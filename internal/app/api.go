package api

import (
	"polonium/internal/pkg/configuration"
)

func setConfiguration(configPath string) {
	configuration.Setup(&configPath)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.json"
	}
	setConfiguration(configPath)
}
