package persistence

import (
	"testing"
	"ui-test-manager/internal/pkg/configuration"
	"ui-test-manager/internal/pkg/model"

	"github.com/google/uuid"
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
		if err := repo.SaveSequence(testSequence); err != nil {
			t.Fatalf("Failed to save sequence: %v", err)
		}

		sequences, err := repo.GetSequences()
		if err != nil {
			t.Fatalf("Failed to get sequences: %v", err)
		}

		if _, ok := (*sequences)[testSequence.Id]; !ok {
			t.Errorf("Saved sequence not found in retrieved sequences")
		}
	})

	t.Run("DeletePrecedent", func(t *testing.T) {
		precedentToDelete := testSequence.Precedents[0]

		if err := repo.DeletePrecedent(testSequence.Id, precedentToDelete); err != nil {
			t.Fatalf("Failed to delete precedent: %v", err)
		}

		sequences, err := repo.GetSequences()
		if err != nil {
			t.Fatalf("Failed to get sequences: %v", err)
		}

		updatedSequence, ok := (*sequences)[testSequence.Id]
		if !ok {
			t.Errorf("Sequence not found after deleting precedent")
		}

		for _, precID := range updatedSequence.Precedents {
			if precID == precedentToDelete {
				t.Errorf("Deleted precedent still found in updated sequence")
			}
		}
	})

	t.Run("DeleteSequence", func(t *testing.T) {
		if err := repo.DeleteSequence(testSequence.Id); err != nil {
			t.Fatalf("Failed to delete sequence: %v", err)
		}

		sequences, err := repo.GetSequences()
		if err != nil {
			t.Fatalf("Failed to get sequences: %v", err)
		}

		if _, ok := (*sequences)[testSequence.Id]; ok {
			t.Errorf("Deleted sequence found in retrieved sequences")
		}
	})

}
