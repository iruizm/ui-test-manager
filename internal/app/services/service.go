package service

import (
	"fmt"
	"strings"
	"ui-test-manager/internal/pkg/model"
	"ui-test-manager/internal/pkg/persistence"

	"github.com/google/uuid"
)

type Service struct {
	sequenceRepo persistence.SequenceRepository
	patternRepo  persistence.PatternRepository
}

func NewService(sequenceRepo persistence.SequenceRepository, patternRepo persistence.PatternRepository) *Service {
	return &Service{
		sequenceRepo: sequenceRepo,
		patternRepo:  patternRepo,
	}
}

func (s *Service) GetSequences() (*map[uuid.UUID]model.Sequence, error) {
	return s.sequenceRepo.GetSequences()
}

func (s *Service) SaveSequence(json model.Sequence) (string, error) {
	zero := uuid.UUID{}
	if json.Id.String() == zero.String() {
		json = *model.NewSequence(json.Name, json.Content)
	}
	return json.Id.String(), s.sequenceRepo.SaveSequence(&json)
}

func (s *Service) DeleteSequence(id uuid.UUID) (string, error) {
	return id.String(), s.sequenceRepo.DeleteSequence(id)
}

func (s *Service) DeletePrecedent(idSequence uuid.UUID, idPrecedent uuid.UUID) (string, error) {
	return idSequence.String() + "|" + idPrecedent.String(), s.sequenceRepo.DeletePrecedent(idSequence, idPrecedent)
}

func (s *Service) OrderSequences() ([]model.Sequence, error) {
	sequences, err := s.sequenceRepo.GetSequences()
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
		return nil, fmt.Errorf("graph contains a cycle can't sort nodes")
	}

	return result, nil
}

func (s *Service) GetPatterns() (*map[uuid.UUID]model.Pattern, error) {
	return s.patternRepo.GetPatterns()
}

func (s *Service) SavePattern(json model.Pattern) (string, error) {
	zero := uuid.UUID{}
	if json.Id.String() == zero.String() {
		json = *model.NewPattern()
	}
	return json.Id.String(), s.patternRepo.SavePattern(&json)
}

func (s *Service) DeletePattern(id uuid.UUID) (string, error) {
	return id.String(), s.patternRepo.DeletePattern(id)
}

func (s *Service) GetTests() ([]model.Sequence, error) {
	o, err := s.OrderSequences()
	if err != nil {
		return nil, err
	}
	patterns, err := s.GetPatterns()
	if err != nil {
		return nil, err
	}

	for k, seq := range o {
		for _, p := range *patterns {
			if p.Replace {
				seq.Content = strings.Replace(seq.Content, p.Before, p.After, -1)
				o[k] = seq
			}
		}
	}
	return o, nil
}
