package api

import (
	"fmt"
	"os"
	"path/filepath"
	"polonium/internal/pkg/configuration"
)

func Run() {
	configPath := ""
	apiDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting API directory:", err)
		os.Exit(1)
	}
	configPath = filepath.Join(apiDir, "configs", "config.json")
	configuration.Setup(&configPath)
	backend := Setup()
	backend.Run("localhost:" + configuration.Config.BackPort)
}
