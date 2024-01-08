package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

var Config *Configuration

type Configuration struct {
	DataPath      string `json:"data_path"`
	ConfigPath    string `json:"config_path"`
	SequencesPath string `json:"sequences_path"`
	PatternsPath  string `json:"patterns_path"`
	LoggingPath   string `json:"logging_path"`
	BackPort      string `json:"back_port"`
	FrontPort     string `json:"front_port"`
}

func Setup(configPath *string) {
	configFile, err := os.ReadFile(*configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(configFile, &Config); err != nil {
		fmt.Println("Error parsing config file:", err)
		os.Exit(1)
	}
}
