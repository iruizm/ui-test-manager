package persistence

import (
	"encoding/json"
	"fmt"
	"os"
	"polonium/internal/pkg/configuration"
	"polonium/internal/pkg/model"
	"sync"

	"github.com/google/uuid"
)

type SequenceMap struct {
	MapData map[uuid.UUID]interface{} `json:"map_data"`
	mu      sync.Mutex
}

var sequenceMap *SequenceMap

func GetSequences() *map[uuid.UUID]interface{} {
	sequences := getSequenceMap()
	return &sequences.MapData

}

func SaveSequence(sequence *model.Sequence) {
	sequences := getSequenceMap()
	sequences.mu.Lock()
	defer sequences.mu.Unlock()
	sequences.MapData[sequence.Id] = sequence
	sequences.saveMap()

}

func DeleteSequence(id *uuid.UUID) {
	sequences := getSequenceMap()
	sequences.mu.Lock()
	defer sequences.mu.Unlock()

	delete(sequences.MapData, *id)
	sequences.saveMap()
}

func getSequenceMap() *SequenceMap {
	if sequenceMap == nil {
		file, errRead := os.ReadFile(configuration.Config.SequencesPath)
		if errRead != nil {
			if os.IsNotExist(errRead) {
				sequenceMap = &SequenceMap{
					MapData: make(map[uuid.UUID]interface{}),
					mu:      sync.Mutex{},
				}

				jsonData, errMarshal := json.Marshal(sequenceMap)
				if errMarshal != nil {
					fmt.Println("Error marshaling initial data:", errMarshal)
					return nil
				}

				if errWrite := os.WriteFile("sequences.json", jsonData, 0644); errWrite != nil {
					fmt.Println("Error writing initial data to file:", errWrite)
					return nil
				}

			} else {
				fmt.Println("Error reading sequence file:", errRead)
				return nil
			}
		} else {
			if err := json.Unmarshal(file, &sequenceMap); err != nil {
				fmt.Println("Error parsing sequence file:", err)
				return nil
			}
		}

	}
	return sequenceMap
}

func (s *SequenceMap) saveMap() {

	jsonData, errMarshal := json.MarshalIndent(&s, "", "  ")
	if errMarshal != nil {
		fmt.Println("Error marshaling initial data:", errMarshal)
		return
	}

	if errWrite := os.WriteFile(configuration.Config.SequencesPath, jsonData, 0644); errWrite != nil {
		fmt.Println("Error writing initial data to file:", errWrite)
		return
	}
}
