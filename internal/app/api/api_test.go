package api_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	controllers "ui-test-manager/internal/app/controllers"
	router "ui-test-manager/internal/app/router"
	services "ui-test-manager/internal/app/services"
	configuration "ui-test-manager/internal/pkg/configuration"
	persistence "ui-test-manager/internal/pkg/persistence"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tempDir := os.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	dummyConfigContent := []byte(`{"dummy": "value"}`)
	dummyConfigFile, err := os.Create(configPath)
	assert.NoError(t, err)
	defer dummyConfigFile.Close()
	_, err = dummyConfigFile.Write(dummyConfigContent)
	assert.NoError(t, err)

	fileConfigurator := &configuration.FileConfigurator{}
	config, err := configuration.SetupConfigurator(fileConfigurator, configPath)
	assert.NoError(t, err)
	configuration.SetConfig(config)

	sequenceRepo := persistence.NewFileSequenceRepository(tempDir, "sequences.json")
	patternRepo := persistence.NewFilePatternRepository(tempDir, "patterns.json")
	service := services.NewService(sequenceRepo, patternRepo)

	controller := controllers.NewController(service)

	backend := router.Setup(controller)

	go backend.Run("localhost:8080")

	req, err := http.NewRequest("GET", "http://localhost:8080/api/health", nil)
	assert.NoError(t, err)

	respRecorder := httptest.NewRecorder()

	backend.ServeHTTP(respRecorder, req)

	assert.Equal(t, http.StatusOK, respRecorder.Code)
}
