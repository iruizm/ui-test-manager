package api

import (
	"fmt"
	"polonium/internal/pkg/model"
	"polonium/internal/pkg/persistence"
)

func GenerateTests() {

}

// Perform topological sorting to order elements by precedence
func OrderSequences() ([]model.Sequence, error) {

	print("TEST")
	sequences, err := persistence.GetSequences()
	if err != nil {
		return nil, err
	}
	queue := make([]model.Sequence, 0)
	for _, seq := range *sequences {
		if len(seq.Precedents) == 0 {
			queue = append(queue, seq)
		}
	}
	var result []model.Sequence
	for !(len(queue) == 0) {

		current := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = make([]model.Sequence, 0)
		}
		result = append(result, current)
		for _, precedent := range current.Precedents {
			if next, ok := (*sequences)[precedent]; ok {
				queue = append(queue, next)
			}

		}
	}
	// Check for cycles (if not all nodes were visited)
	if len(result) != len(*sequences) {
		return nil, fmt.Errorf("graph contains a cycle")
	}
	print(result)
	return result, nil

}
