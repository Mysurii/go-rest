package main

import (
	"fmt"
	"go-rest/internal/models"
	"go-rest/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type projectConfig struct {
	Name   string
	DBType string
}

type TemplateInfo struct {
	FilePath 	 string
	TemplatePath string
}



// Lipgloss styles for the views
var (
	bannerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e879f9")).
			Bold(true).
			Padding(1, 2).
			Align(lipgloss.Center)
)





// Main function where Bubble Tea's tea.NewProgram is initialized
func main() {

	// Banner for go-rest title
	banner := `
>>===========================================================<<
||                                                           ||
||  ██████╗  ██████╗       ██████╗ ███████╗███████╗████████╗ ||
|| ██╔════╝ ██╔═══██╗      ██╔══██╗██╔════╝██╔════╝╚══██╔══╝ ||
|| ██║  ███╗██║   ██║█████╗██████╔╝█████╗  ███████╗   ██║    ||
|| ██║   ██║██║   ██║╚════╝██╔══██╗██╔══╝  ╚════██║   ██║    ||
|| ╚██████╔╝╚██████╔╝      ██║  ██║███████╗███████║   ██║    ||
||  ╚═════╝  ╚═════╝       ╚═╝  ╚═╝╚══════╝╚══════╝   ╚═╝    ||
||                                                           ||
>>===========================================================<<

GO-REST
version: 1.2.1

	 `
	// Format the banner
	bannerOutput := bannerStyle.Render(banner)

	println(bannerOutput)


	// Initialize the model for entering the project name
	projectInputModel := models.InitializeProjectModel("Name of your project.")

	// Run the program for project name input
	finalProjectModel, err := tea.NewProgram(projectInputModel).Run()
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}

	// Retrieve the project name from the final model
	projectModel, ok := finalProjectModel.(models.ProjectModel)  // Assuming the type is models.ProjectModel
	if !ok {
		log.Fatalf("unexpected model type")
	}
	


	// Initialize the model for selecting a database
	databaseSelectionModel := models.InitialDriverModel("Select the database you want to use for your server:")

	// Run the program for database selection
	finalDatabaseModel, err := tea.NewProgram(databaseSelectionModel).Run()
	if err != nil {
		log.Fatalf("could not start program: %s", err)
	}

	// Retrieve the selected database from the final model
	driverModel, ok := finalDatabaseModel.(models.DriverModel)  // Assuming the type is models.DriverModel
	if !ok {
		log.Fatalf("unexpected model type")
	}


	// Now you can print the results
	println("ProjectName: " + projectModel.GetName())
	println("Chosen DB: " + driverModel.GetDriver())


	baseDir := filepath.Join(".", projectModel.GetName())
	dirs := []string{
		filepath.Join(baseDir, "cmd", "api"),
		filepath.Join(baseDir, "internal", "server"),
		filepath.Join(baseDir, "internal", "database"),
	}

	templates := []TemplateInfo{
		{filepath.Join(baseDir, "cmd", "api", "main.go"), "internal/templates/cmd.api.main.go.templ"},
		{filepath.Join(baseDir, "internal", "server", "server.go"), "internal/templates/server.server.go.templ"},
		{filepath.Join(baseDir, "internal", "server", "routes.go"), "internal/templates/server.routes.go.templ"},
		{filepath.Join(baseDir, "internal", "database", "database.go"), fmt.Sprintf("internal/templates/db.%s.go.templ", strings.ToLower(driverModel.GetDriver()))},
		{filepath.Join(baseDir, ".env"), "internal/templates/env.templ"},
	}

	// Parse all templates first to avoid partial creation
	for _, tmplInfo := range templates {
		err := utils.ValidateTemplate(tmplInfo.TemplatePath)
		if err != nil {
			log.Fatalf("Template validation failed: %v", err)
		}
	}

		// Create directories
		for _, dir := range dirs {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				log.Fatalf("Failed to create directory %s: %v", dir, err)
			}
		}
	
		// Channel to track file generation status
		statusChan := make(chan string)

		templData := struct {
			Project string
			Driver string
		}{
			Project: projectModel.GetName(), // Assign the value you need
			Driver: strings.ToLower(driverModel.GetDriver()),
		}
	
		// Generate files using goroutines
		for _, tmplInfo := range templates {
			go utils.GenerateFileFromTemplate(tmplInfo.FilePath, tmplInfo.TemplatePath, templData, statusChan)
		}
	
		// Receive results from the channel for each template
		for i := 0; i < len(templates); i++ {
			status := <-statusChan
			fmt.Println(status)
		}

		err = utils.RunGoModInit(projectModel.GetName())
		if err != nil {
			println(err.Error())
		}

		
		err = utils.RunGoModTidy(projectModel.GetName())
		if err != nil {
			println(err.Error())
		}
	
}
