package api

import (
	"fmt"
	"os"
	"path/filepath"
	router "selenium-manager/internal/app/router"
	configuration "selenium-manager/internal/pkg/configuration"
)

func Run() {
	configPath := ""
	exeDir, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting API directory:", err)
		os.Exit(1)
	}
	configPath = filepath.Join(filepath.Dir(filepath.Dir(exeDir)), "configs", "config.json")
	configuration.Setup(&configPath)
	backend := router.Setup()
	backend.Run("localhost:" + configuration.Config.BackPort)
}
