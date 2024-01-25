package configuration

import (
	"os"
	"testing"
)

func TestFileConfigurator_ReadConfig(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_config.json")
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	testConfigData := `{
		"root_path": "/path/to/root",
		"data_path": "/path/to/data",
		"config_path": "/path/to/config",
		"sequences_path": "/path/to/sequences",
		"patterns_path": "/path/to/patterns",
		"logging_path": "/path/to/logging",
		"back_port": "8080",
		"front_port": "3000"
	}`
	err = os.WriteFile(tmpFile.Name(), []byte(testConfigData), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config data to temporary file: %v", err)
	}

	configurator := &FileConfigurator{}

	config, err := configurator.ReadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read config from file: %v", err)
	}

	expectedConfig := &Configuration{
		RootPath:      "/path/to/root",
		DataPath:      "/path/to/data",
		ConfigPath:    "/path/to/config",
		SequencesPath: "/path/to/sequences",
		PatternsPath:  "/path/to/patterns",
		LoggingPath:   "/path/to/logging",
		BackPort:      "8080",
		FrontPort:     "3000",
	}

	if *config != *expectedConfig {
		t.Errorf("Read config does not match expected config:\nRead: %+v\nExpected: %+v", config, expectedConfig)
	}
}
