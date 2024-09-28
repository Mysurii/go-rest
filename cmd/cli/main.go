package main

import (
	"fmt"
	"go-rest/internal/config"
	"go-rest/internal/models"
	"go-rest/pkg/utils"
	"log"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	p = &models.Program{}
)


func printBanner() {
	fmt.Println(config.Banner)
}

func promptProjectName() error {
	model := models.InitializeProjectModel("Name of your project.", p)
	tProgram := tea.NewProgram(model)

	if _, err := tProgram.Run(); err != nil {
		return err
	}

	p.ExitCLI(tProgram)

	return nil
}

func promtDriverSelection() error {
	model := models.InitialDriverModel("Select the database you want to use for your server:", p)
	tProgram := tea.NewProgram(model)
	if _, err := tProgram.Run(); err != nil {
		return err
	}

	p.ExitCLI(tProgram)

	return nil
}




// Main function where Bubble Tea's tea.NewProgram is initialized
func main() {	
	printBanner()

	err := promptProjectName()
	if err != nil {
		log.Fatalf("Error retrieving project name: %v", err)
	}


	err = promtDriverSelection()
	if err != nil {
		log.Fatalf("Error retrieving database: %v", err)
	}

	println("Project: ", p.Project)
	println("Driver:", p.GetDriver())

	baseDir := filepath.Join(".", p.Project)
	config.CreateProjectStructure(baseDir)

	templData := models.TemplateData{
		Project: p.Project,
		Driver:  p.GetDriver(),
	}

	templates := config.PrepareTemplates(baseDir, p.GetDriver())

	for _, t := range templates {
		if err := utils.ValidateTemplate(t.TemplatePath); err != nil {
			panic(err)
		}
	}



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
	if err := config.InitializeGoMod(p.Project); err != nil {
		log.Println(err)
	}
	
}
