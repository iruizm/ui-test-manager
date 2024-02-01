package persistence

import (
	"testing"
	"ui-test-manager/internal/pkg/configuration"
	"ui-test-manager/internal/pkg/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPatternRepository(t *testing.T) {
	tmpDir := t.TempDir()
	configuration.Config.PatternsPath = "test_patterns.json"
	configuration.Config.DataPath = tmpDir

	testPattern := &model.Pattern{
		Name:   "Test Pattern",
		Id:     uuid.New(),
		Before: "Before Test",
		After:  "After Test",
	}

	repo := NewFilePatternRepository(configuration.Config.PatternsPath, configuration.Config.DataPath)

	t.Run("SavePattern and GetPatterns", func(t *testing.T) {
		err := repo.SavePattern(testPattern)
		assert.NoError(t, err)

		patterns, err := repo.GetPatterns()
		assert.NoError(t, err)

		assert.NotNil(t, patterns)
		assert.Contains(t, *patterns, testPattern.Id, "Saved pattern not found in retrieved patterns")
	})

	t.Run("DeletePattern", func(t *testing.T) {
		err := repo.DeletePattern(testPattern.Id)
		assert.NoError(t, err)

		patterns, err := repo.GetPatterns()
		assert.NoError(t, err)

		assert.NotContains(t, *patterns, testPattern.Id, "Deleted pattern found in retrieved patterns")
	})
}
