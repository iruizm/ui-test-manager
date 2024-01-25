package api

import (
	"fmt"
	"os"
	"path/filepath"
	controllers "ui-test-manager/internal/app/controllers"
	router "ui-test-manager/internal/app/router"
	services "ui-test-manager/internal/app/services"
	configuration "ui-test-manager/internal/pkg/configuration"
	persistence "ui-test-manager/internal/pkg/persistence"
)

func Run() {
	configPath := ""
	exeDir, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting API directory:", err)
		os.Exit(1)
	}
	configPath = filepath.Join(filepath.Dir(filepath.Dir(exeDir)), "configs", "config.json")

	fileConfigurator := &configuration.FileConfigurator{}
	config, err := configuration.SetupConfigurator(fileConfigurator, configPath)
	if err != nil {
		fmt.Println("Error setting up configurator:", err)
		os.Exit(1)
	}

	configuration.SetConfig(config)

	sequenceRepo := persistence.GetSequenceRepository()
	patternRepo := persistence.GetPatternRepository()

	service := services.NewService(sequenceRepo, patternRepo)

	controller := controllers.NewController(service)

	backend := router.Setup(controller)
	fmt.Println("Server running...")
	backend.Run("localhost:" + configuration.Config.BackPort)

}
