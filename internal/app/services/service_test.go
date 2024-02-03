package service_test

import (
	"testing"
	service "ui-test-manager/internal/app/services"
	"ui-test-manager/internal/pkg/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockSequenceRepository struct {
	sequences map[uuid.UUID]model.Sequence
}

type mockPatternRepository struct {
	patterns map[uuid.UUID]model.Pattern
}

func (m *mockSequenceRepository) GetSequences() (*map[uuid.UUID]model.Sequence, error) {
	return &m.sequences, nil
}

func (m *mockSequenceRepository) SaveSequence(seq *model.Sequence) error {
	m.sequences[seq.Id] = *seq
	return nil
}

func (m *mockSequenceRepository) DeleteSequence(id uuid.UUID) error {
	delete(m.sequences, id)
	return nil
}

func (m *mockSequenceRepository) DeletePrecedent(idSequence uuid.UUID, idPrecedent uuid.UUID) error {
	seq := m.sequences[idSequence]
	for i := 0; i < len(seq.Precedents); i++ {
		if seq.Precedents[i] == idPrecedent {
			seq.Precedents = append(seq.Precedents[:i], seq.Precedents[i+1:]...)
		}
	}
	m.sequences[idSequence] = seq
	return nil
}

func (m *mockPatternRepository) GetPatterns() (*map[uuid.UUID]model.Pattern, error) {
	return &m.patterns, nil
}

func (m *mockPatternRepository) SavePattern(pattern *model.Pattern) error {
	m.patterns[pattern.Id] = *pattern
	return nil
}

func (m *mockPatternRepository) DeletePattern(id uuid.UUID) error {
	delete(m.patterns, id)
	return nil
}

func TestService_GetSequences(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	seq1 := model.Sequence{Id: uuid.New(), Name: "Sequence 1", Content: "Content 1"}
	seq2 := model.Sequence{Id: uuid.New(), Name: "Sequence 2", Content: "Content 2"}

	mockSeqRepo.sequences[seq1.Id] = seq1
	mockSeqRepo.sequences[seq2.Id] = seq2

	result, err := s.GetSequences()
	assert.NoError(t, err)
	assert.Len(t, *result, 2)
	assert.Contains(t, *result, seq1.Id)
	assert.Contains(t, *result, seq2.Id)
}

func TestService_SaveSequence(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	seq := model.Sequence{Id: uuid.New(), Name: "Sequence 1", Content: "Content 1"}

	resultID, err := s.SaveSequence(seq)
	assert.NoError(t, err)
	assert.Equal(t, seq.Id.String(), resultID)
	assert.Contains(t, mockSeqRepo.sequences, seq.Id)
}

func TestService_DeleteSequence(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	seq := model.Sequence{Id: uuid.New(), Name: "Sequence 1", Content: "Content 1"}
	mockSeqRepo.sequences[seq.Id] = seq

	resultID, err := s.DeleteSequence(seq.Id)
	assert.NoError(t, err)
	assert.Equal(t, seq.Id.String(), resultID)
	assert.NotContains(t, mockSeqRepo.sequences, seq.Id)
}

func TestService_DeletePrecedent(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	seqID := uuid.New()
	precedentID := uuid.New()
	mockSeqRepo.sequences[seqID] = model.Sequence{Id: seqID, Name: "Sequence 1", Content: "Content 1", Precedents: []uuid.UUID{precedentID}}

	result, err := s.DeletePrecedent(seqID, precedentID)
	assert.NoError(t, err)
	assert.Equal(t, seqID.String()+"|"+precedentID.String(), result)
	assert.NotContains(t, mockSeqRepo.sequences[seqID].Precedents, precedentID)
}

func TestService_GetPatterns(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	pattern1 := model.Pattern{Id: uuid.New(), Before: "Before 1", After: "After 1"}
	pattern2 := model.Pattern{Id: uuid.New(), Before: "Before 2", After: "After 2"}

	mockPatternRepo.patterns[pattern1.Id] = pattern1
	mockPatternRepo.patterns[pattern2.Id] = pattern2

	result, err := s.GetPatterns()
	assert.NoError(t, err)
	assert.Len(t, *result, 2)
	assert.Contains(t, *result, pattern1.Id)
	assert.Contains(t, *result, pattern2.Id)
}

func TestService_SavePattern(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	pattern := model.Pattern{Id: uuid.New(), Before: "Before 1", After: "After 1"}

	resultID, err := s.SavePattern(pattern)
	assert.NoError(t, err)
	assert.Equal(t, pattern.Id.String(), resultID)
	assert.Contains(t, mockPatternRepo.patterns, pattern.Id)
}

func TestService_DeletePattern(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	pattern := model.Pattern{Id: uuid.New(), Before: "Before 1", After: "After 1"}
	mockPatternRepo.patterns[pattern.Id] = pattern

	resultID, err := s.DeletePattern(pattern.Id)
	assert.NoError(t, err)
	assert.Equal(t, pattern.Id.String(), resultID)
	assert.NotContains(t, mockPatternRepo.patterns, pattern.Id)
}

func TestService_GetTests(t *testing.T) {
	mockSeqRepo := &mockSequenceRepository{
		sequences: make(map[uuid.UUID]model.Sequence),
	}
	mockPatternRepo := &mockPatternRepository{
		patterns: make(map[uuid.UUID]model.Pattern),
	}
	s := service.NewService(mockSeqRepo, mockPatternRepo)

	seq1 := model.Sequence{Id: uuid.New(), Name: "Sequence 1", Content: "Content 1"}
	seq2 := model.Sequence{Id: uuid.New(), Name: "Sequence 2", Content: "Content 2"}

	mockSeqRepo.sequences[seq1.Id] = seq1
	mockSeqRepo.sequences[seq2.Id] = seq2

	pattern := model.Pattern{Id: uuid.New(), Before: "Content 1", After: "Content 2"}
	mockPatternRepo.patterns[pattern.Id] = pattern

	result, err := s.GetTests()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, result[0].Content, "Content 2")
	assert.Equal(t, result[1].Content, "Content 2")
}
