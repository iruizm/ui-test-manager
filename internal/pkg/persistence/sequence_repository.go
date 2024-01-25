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

type SequenceRepository interface {
	GetSequences() (*map[uuid.UUID]model.Sequence, error)
	SaveSequence(sequence *model.Sequence) error
	DeleteSequence(id uuid.UUID) error
	DeletePrecedent(idSequence uuid.UUID, idPrecedent uuid.UUID) error
}

type fileSequenceRepository struct {
	filePath string
	dataPath string
	mu       sync.Mutex
}

func NewFileSequenceRepository(filePath, dataPath string) SequenceRepository {
	return &fileSequenceRepository{
		filePath: filePath,
		dataPath: dataPath,
	}
}

func (r *fileSequenceRepository) GetSequences() (*map[uuid.UUID]model.Sequence, error) {
	filePath := filepath.Join(r.dataPath, r.filePath)
	file, errRead := os.ReadFile(filePath)
	if errRead != nil {
		if os.IsNotExist(errRead) {
			emptyMap := make(map[uuid.UUID]model.Sequence)
			if errWrite := r.saveMap(&emptyMap); errWrite != nil {
				return nil, errWrite
			}
			return &emptyMap, nil
		}
		return nil, errRead
	}

	var sequenceMap map[uuid.UUID]model.Sequence
	if err := json.Unmarshal(file, &sequenceMap); err != nil {
		return nil, err
	}

	return &sequenceMap, nil
}
func (r *fileSequenceRepository) SaveSequence(sequence *model.Sequence) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sequenceMap, err := r.GetSequences()
	if err != nil {
		return err
	}

	(*sequenceMap)[sequence.Id] = *sequence

	return r.saveMap(sequenceMap)
}

func (r *fileSequenceRepository) DeleteSequence(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sequenceMap, err := r.GetSequences()
	if err != nil {
		return err
	}

	delete(*sequenceMap, id)

	return r.saveMap(sequenceMap)
}

func (r *fileSequenceRepository) DeletePrecedent(idSequence uuid.UUID, idPrecedent uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sequenceMap, err := r.GetSequences()
	if err != nil {
		return err
	}

	sequence, found := (*sequenceMap)[idSequence]
	if !found {
		return nil
	}

	updatedPrecedents := make([]uuid.UUID, 0)
	for _, precID := range sequence.Precedents {
		if precID != idPrecedent {
			updatedPrecedents = append(updatedPrecedents, precID)
		}
	}

	sequence.Precedents = updatedPrecedents
	(*sequenceMap)[idSequence] = sequence

	return r.saveMap(sequenceMap)
}

func (r *fileSequenceRepository) saveMap(sequenceMap *map[uuid.UUID]model.Sequence) error {
	if _, err := os.Stat(r.dataPath); os.IsNotExist(err) {
		if errDir := os.MkdirAll(r.dataPath, os.ModePerm); errDir != nil {
			return errDir
		}
	}

	jsonData, errMarshal := json.MarshalIndent(sequenceMap, "", "  ")
	if errMarshal != nil {
		return errMarshal
	}

	if errWrite := os.WriteFile(filepath.Join(r.dataPath, r.filePath), jsonData, 0644); errWrite != nil {
		return errWrite
	}

	return nil
}

func GetSequenceRepository() SequenceRepository {
	return NewFileSequenceRepository(configuration.Config.SequencesPath, configuration.Config.DataPath)
}
