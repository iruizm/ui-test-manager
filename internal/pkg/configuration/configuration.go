package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

var Config Configuration

type Configurator interface {
	ReadConfig(configPath string) (*Configuration, error)
}

func SetConfig(config *Configuration) {
	Config = *config
}

type Configuration struct {
	RootPath      string `json:"root_path"`
	DataPath      string `json:"data_path"`
	ConfigPath    string `json:"config_path"`
	SequencesPath string `json:"sequences_path"`
	PatternsPath  string `json:"patterns_path"`
	LoggingPath   string `json:"logging_path"`
	BackPort      string `json:"back_port"`
	FrontPort     string `json:"front_port"`
}

func SetupConfigurator(configurator Configurator, configPath string) (*Configuration, error) {
	return configurator.ReadConfig(configPath)
}

type FileConfigurator struct{}

func (c *FileConfigurator) ReadConfig(configPath string) (*Configuration, error) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Configuration
	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}
