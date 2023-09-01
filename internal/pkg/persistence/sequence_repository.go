package persistence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"polonium/internal/pkg/configuration"
	"polonium/internal/pkg/model"
	"sync"

	"github.com/google/uuid"
)

type SequenceMap struct {
	MapData map[uuid.UUID]model.Sequence `json:"map_data"`
	mu      sync.Mutex
}

var sequenceMap *SequenceMap

func GetSequences() (*map[uuid.UUID]model.Sequence, error) {
	sequences, err := getSequenceMap()
	if err != nil {
		return nil, err
	}
	return &sequences.MapData, nil
}

func SaveSequence(sequence *model.Sequence) error {
	sequences, err := getSequenceMap()
	if err != nil {
		return err
	}
	sequences.mu.Lock()
	defer sequences.mu.Unlock()
	sequences.MapData[sequence.Id] = *sequence
	err = sequences.saveMap()
	if err != nil {
		return err
	}
	return nil
}

func DeleteSequence(id *uuid.UUID) error {
	sequences, err := getSequenceMap()
	if err != nil {
		return err
	}
	sequences.mu.Lock()
	defer sequences.mu.Unlock()

	for key, sequence := range sequences.MapData {
		updatedPrecedents := make([]uuid.UUID, 0)
		for _, precID := range sequence.Precedents {
			if precID != *id {
				updatedPrecedents = append(updatedPrecedents, precID)
			}
		}
		sequence.Precedents = updatedPrecedents
		sequences.MapData[key] = sequence
	}
	delete(sequences.MapData, *id)
	err = sequences.saveMap()
	if err != nil {
		return err
	}
	return nil
}
func DeletePrecedent(idSequence *uuid.UUID, idPrecedent *uuid.UUID) error {
	sequences, err := getSequenceMap()
	if err != nil {
		return err
	}
	sequences.mu.Lock()
	defer sequences.mu.Unlock()

	sequence := sequences.MapData[*idSequence]
	updatedPrecedents := make([]uuid.UUID, 0)
	for _, precID := range sequence.Precedents {
		if precID != *idPrecedent {
			updatedPrecedents = append(updatedPrecedents, precID)
		}
	}
	sequence.Precedents = updatedPrecedents
	sequences.MapData[*idSequence] = sequence

	err = sequences.saveMap()
	if err != nil {
		return err
	}
	return nil
}

func getSequenceMap() (*SequenceMap, error) {
	if sequenceMap == nil {
		file, errRead := os.ReadFile(filepath.Join(configuration.Config.DataPath, configuration.Config.SequencesPath))
		if errRead != nil {
			if os.IsNotExist(errRead) {
				sequenceMap = &SequenceMap{
					MapData: make(map[uuid.UUID]model.Sequence),
					mu:      sync.Mutex{},
				}

				jsonData, errMarshal := json.Marshal(sequenceMap)
				if errMarshal != nil {
					return nil, errMarshal
				}

				if errWrite := os.WriteFile(filepath.Join(configuration.Config.DataPath, configuration.Config.SequencesPath), jsonData, 0644); errWrite != nil {
					return nil, errWrite
				}

			} else {
				return nil, errRead
			}
		} else {
			if err := json.Unmarshal(file, &sequenceMap); err != nil {
				return nil, err
			}
		}

	}
	return sequenceMap, nil
}

func (s *SequenceMap) saveMap() error {

	jsonData, errMarshal := json.MarshalIndent(&s, "", "  ")
	if errMarshal != nil {
		return errMarshal
	}
	if errWrite := os.WriteFile(filepath.Join(configuration.Config.DataPath, configuration.Config.SequencesPath), jsonData, 0644); errWrite != nil {
		return errWrite
	}
	return nil
}
