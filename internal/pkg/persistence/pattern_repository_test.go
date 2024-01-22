package persistence

import (
	"testing"
	"ui-test-manager/internal/pkg/configuration"
	"ui-test-manager/internal/pkg/model"

	"github.com/google/uuid"
)

func TestPatternRepository(t *testing.T) {
	tmpDir := t.TempDir()
	configuration.Config.PatternsPath = "test_patterns.json"
	configuration.Config.DataPath = tmpDir

	testPattern := &model.Pattern{
		Name:    "Test Pattern",
		Id:      uuid.New(),
		Before:  "Before Test",
		After:   "After Test",
		Replace: true,
	}

	repo := NewFilePatternRepository(configuration.Config.PatternsPath, configuration.Config.DataPath)

	t.Run("SavePattern and GetPatterns", func(t *testing.T) {
		if err := repo.SavePattern(testPattern); err != nil {
			t.Fatalf("Failed to save pattern: %v", err)
		}

		patterns, err := repo.GetPatterns()
		if err != nil {
			t.Fatalf("Failed to get patterns: %v", err)
		}

		if _, ok := (*patterns)[testPattern.Id]; !ok {
			t.Errorf("Saved pattern not found in retrieved patterns")
		}
	})

	t.Run("DeletePattern", func(t *testing.T) {
		if err := repo.DeletePattern(testPattern.Id); err != nil {
			t.Fatalf("Failed to delete pattern: %v", err)
		}

		patterns, err := repo.GetPatterns()
		if err != nil {
			t.Fatalf("Failed to get patterns: %v", err)
		}

		if _, ok := (*patterns)[testPattern.Id]; ok {
			t.Errorf("Deleted pattern found in retrieved patterns")
		}
	})
}
