package service

import (
	"fmt"
	"selenium-manager/internal/pkg/model"
	"selenium-manager/internal/pkg/persistence"

	"github.com/google/uuid"
)

func GetSequences() (*map[uuid.UUID]model.Sequence, error) {
	return persistence.GetSequences()
}

func SaveSequence(json model.Sequence) (string, error) {
	zero := uuid.UUID{}
	if json.Id.String() == zero.String() {
		json = *model.NewSequence(json.Name, json.Content)
	}
	return json.Id.String(), persistence.SaveSequence(&json)
}

func DeleteSequence(id uuid.UUID) (string, error) {
	return id.String(), persistence.DeleteSequence(&id)
}

func DeletePrecedent(idSequence uuid.UUID, idPrecedent uuid.UUID) (string, error) {
	return idSequence.String() + "|" + idPrecedent.String(), persistence.DeletePrecedent(&idSequence, &idPrecedent)
}

func OrderSequences() ([]model.Sequence, error) {
	sequences, err := persistence.GetSequences()
	if err != nil {
		return nil, err
	}

	seqMap := make(map[uuid.UUID]model.Sequence)
	for _, seq := range *sequences {
		seqMap[seq.Id] = seq
	}

	inDegree := make(map[uuid.UUID]int)
	for _, seq := range *sequences {
		for _, prec := range seq.Precedents {
			inDegree[prec]++
		}
	}

	stack := make([]model.Sequence, 0)
	for _, seq := range *sequences {
		if inDegree[seq.Id] == 0 {
			stack = append(stack, seq)
		}
	}

	result := make([]model.Sequence, 0)

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]model.Sequence{current}, result...)

		for _, prec := range current.Precedents {
			inDegree[prec]--
			if inDegree[prec] == 0 {
				stack = append(stack, seqMap[prec])
			}
		}
	}

	if len(result) != len(*sequences) {
		return nil, fmt.Errorf("graph contains a cycle")
	}

	return result, nil
}

func ReplacePattern(*model.Sequence) {

}
