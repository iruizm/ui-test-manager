package persistence

import (
	"encoding/json"
	"fmt"
	"os"
	configuration "polonium/internal/pkg/configuration"
	"polonium/internal/pkg/model"
	"sync"
)

type SequenceMap struct {
	MapData map[string]interface{} `json:"map_data"`
	mu      sync.Mutex
}

var sequenceMap *SequenceMap

func getSequenceMap() *SequenceMap {
	if sequenceMap == nil {
		file, errRead := os.ReadFile(configuration.Config.SequencesPath)
		if errRead != nil {
			if os.IsNotExist(errRead) {
				sequenceMap = &SequenceMap{}

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

func (s *SequenceMap) addSequence(sequence *model.Sequence) error {
	//sequences := getSequenceMap()
	return nil
}
