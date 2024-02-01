package persistence

import (
	"testing"
	"ui-test-manager/internal/pkg/configuration"
	"ui-test-manager/internal/pkg/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSequenceRepository(t *testing.T) {
	tmpDir := t.TempDir()
	configuration.Config.SequencesPath = "test_sequences.json"
	configuration.Config.DataPath = tmpDir

	testSequence := &model.Sequence{
		Name:       "Test Sequence",
		Id:         uuid.New(),
		Precedents: []uuid.UUID{uuid.New(), uuid.New()},
	}

	repo := NewFileSequenceRepository(configuration.Config.SequencesPath, configuration.Config.DataPath)

	t.Run("SaveSequence and GetSequences", func(t *testing.T) {
		err := repo.SaveSequence(testSequence)
		assert.NoError(t, err)

		sequences, err := repo.GetSequences()
		assert.NoError(t, err)

		assert.NotNil(t, sequences)
		assert.Contains(t, *sequences, testSequence.Id, "Saved sequence not found in retrieved sequences")
	})

	t.Run("DeletePrecedent", func(t *testing.T) {
		precedentToDelete := testSequence.Precedents[0]

		err := repo.DeletePrecedent(testSequence.Id, precedentToDelete)
		assert.NoError(t, err)

		sequences, err := repo.GetSequences()
		assert.NoError(t, err)

		updatedSequence, ok := (*sequences)[testSequence.Id]
		assert.True(t, ok, "Sequence not found after deleting precedent")

		for _, precID := range updatedSequence.Precedents {
			assert.NotEqual(t, precID, precedentToDelete, "Deleted precedent still found in updated sequence")
		}
	})

	t.Run("DeleteSequence", func(t *testing.T) {
		err := repo.DeleteSequence(testSequence.Id)
		assert.NoError(t, err)

		sequences, err := repo.GetSequences()
		assert.NoError(t, err)

		assert.NotContains(t, *sequences, testSequence.Id, "Deleted sequence found in retrieved sequences")
	})
}
