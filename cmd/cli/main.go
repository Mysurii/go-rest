package main

import (
	"fmt"
	"go-rest/internal/config"
	"go-rest/internal/models"
	"go-rest/pkg/utils"
	"log"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)


func printBanner() {
	fmt.Println(config.Banner)
}

func promptProjectName() (models.ProjectModel, error) {
	projectInputModel := models.InitializeProjectModel("Name of your project.")
	finalProjectModel, err := tea.NewProgram(projectInputModel).Run()
	if err != nil {
		return models.ProjectModel{}, err
	}

	projectModel, ok := finalProjectModel.(models.ProjectModel)
	if !ok {
		return models.ProjectModel{}, fmt.Errorf("unexpected model type")
	}

	return projectModel, nil
}

func promptDatabaseSelection() (models.DriverModel, error) {
	databaseSelectionModel := models.InitialDriverModel("Select the database you want to use for your server:")
	finalDatabaseModel, err := tea.NewProgram(databaseSelectionModel).Run()
	if err != nil {
		return models.DriverModel{}, err
	}

	driverModel, ok := finalDatabaseModel.(models.DriverModel)
	if !ok {
		return models.DriverModel{}, fmt.Errorf("unexpected model type")
	}

	return driverModel, nil
}




// Main function where Bubble Tea's tea.NewProgram is initialized
func main() {	
	printBanner()

	projectModel, err := promptProjectName()
	if err != nil {
		log.Fatalf("Error retrieving project name: %v", err)
	}


	driverModel, err := promptDatabaseSelection()
	if err != nil {
		log.Fatalf("Error retrieving database: %v", err)
	}

	baseDir := filepath.Join(".", projectModel.GetName())
	config.CreateProjectStructure(baseDir)

	templData := models.TemplateData{
		Project: projectModel.GetName(),
		Driver:  strings.ToLower(driverModel.GetDriver()),
	}

	templates := config.PrepareTemplates(baseDir, driverModel.GetDriver())

	// Generate files in parallel
	statusChan := make(chan string)
	for _, tmplInfo := range templates {
		go utils.GenerateFileFromTemplate(tmplInfo.FilePath, tmplInfo.TemplatePath, templData, statusChan)
	}

	// Receive and print results
	for i := 0; i < len(templates); i++ {
		fmt.Println(<-statusChan)
	}

	// Initialize Go module and tidy up
	if err := config.InitializeGoMod(projectModel.GetName()); err != nil {
		log.Println(err)
	}
	
}
