package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

var Config *Configuration

type Configuration struct {
	SequencesPath string `json:"sequences_path"`
	LoggingPath   string `json:"logging_path"`
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
	fmt.Println("Start")
}
