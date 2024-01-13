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

type PatternMap struct {
	MapData map[uuid.UUID]model.Pattern `json:"map_data"`
	mu      sync.Mutex
}

var patternMap *PatternMap

func GetPatterns() (*map[uuid.UUID]model.Pattern, error) {
	patterns, err := getPatternMap()
	if err != nil {
		return nil, err
	}
	return &patterns.MapData, nil
}

func SavePattern(pattern *model.Pattern) error {
	patterns, err := getPatternMap()
	if err != nil {
		return err
	}
	patterns.mu.Lock()
	defer patterns.mu.Unlock()
	patterns.MapData[pattern.Id] = *pattern
	err = patterns.saveMap()
	if err != nil {
		return err
	}
	return nil
}

func DeletePattern(id *uuid.UUID) error {
	patterns, err := getPatternMap()
	if err != nil {
		return err
	}
	patterns.mu.Lock()
	defer patterns.mu.Unlock()
	delete(patterns.MapData, *id)
	err = patterns.saveMap()
	if err != nil {
		return err
	}
	return nil
}
func getPatternMap() (*PatternMap, error) {
	if patternMap == nil {
		file, errRead := os.ReadFile(filepath.Join(configuration.Config.DataPath, configuration.Config.PatternsPath))
		if errRead != nil {
			if os.IsNotExist(errRead) {
				patternMap = &PatternMap{
					MapData: make(map[uuid.UUID]model.Pattern),
					mu:      sync.Mutex{},
				}

				jsonData, errMarshal := json.Marshal(patternMap)
				if errMarshal != nil {
					return nil, errMarshal
				}

				if errWrite := os.WriteFile(filepath.Join(configuration.Config.DataPath, configuration.Config.PatternsPath), jsonData, 0644); errWrite != nil {
					return nil, errWrite
				}

			} else {
				return nil, errRead
			}
		} else {
			if err := json.Unmarshal(file, &patternMap); err != nil {
				return nil, err
			}
		}

	}
	return patternMap, nil
}

func (s *PatternMap) saveMap() error {

	jsonData, errMarshal := json.MarshalIndent(&s, "", "  ")
	if errMarshal != nil {
		return errMarshal
	}
	if errWrite := os.WriteFile(filepath.Join(configuration.Config.DataPath, configuration.Config.PatternsPath), jsonData, 0644); errWrite != nil {
		return errWrite
	}
	return nil
}
