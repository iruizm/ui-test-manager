package persistence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"ui-test-manager/internal/pkg/configuration"
	"ui-test-manager/internal/pkg/model"

	"github.com/google/uuid"
)

type PatternRepository interface {
	GetPatterns() (*map[uuid.UUID]model.Pattern, error)
	SavePattern(pattern *model.Pattern) error
	DeletePattern(id uuid.UUID) error
}

type FilePatternRepository struct {
	filePath string
	dataPath string
	mu       sync.Mutex
}

func NewFilePatternRepository(filePath, dataPath string) *FilePatternRepository {
	return &FilePatternRepository{
		filePath: filePath,
		dataPath: dataPath,
	}
}

func (r *FilePatternRepository) GetPatterns() (*map[uuid.UUID]model.Pattern, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, errRead := os.ReadFile(filepath.Join(r.dataPath, r.filePath))
	if errRead != nil {
		if os.IsNotExist(errRead) {
			patternMap := make(map[uuid.UUID]model.Pattern)
			return &patternMap, nil
		}
		return nil, errRead
	}

	var patternMap map[uuid.UUID]model.Pattern
	if err := json.Unmarshal(file, &patternMap); err != nil {
		return nil, err
	}

	return &patternMap, nil
}

func (r *FilePatternRepository) SavePattern(pattern *model.Pattern) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	patternMap, err := r.GetPatterns()
	if err != nil {
		return err
	}

	(*patternMap)[pattern.Id] = *pattern

	return r.saveMap(patternMap)
}

func (r *FilePatternRepository) DeletePattern(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	patternMap, err := r.GetPatterns()
	if err != nil {
		return err
	}

	delete(*patternMap, id)

	return r.saveMap(patternMap)
}

func (r *FilePatternRepository) saveMap(patternMap *map[uuid.UUID]model.Pattern) error {
	jsonData, errMarshal := json.MarshalIndent(patternMap, "", "  ")
	if errMarshal != nil {
		return errMarshal
	}

	if errWrite := os.WriteFile(filepath.Join(r.dataPath, r.filePath), jsonData, 0644); errWrite != nil {
		return errWrite
	}

	return nil
}

func GetPatternRepository() PatternRepository {
	return NewFilePatternRepository(configuration.Config.PatternsPath, configuration.Config.DataPath)
}
