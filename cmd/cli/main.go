package main

import (
	"fmt"
	"go-rest/internal/config"
	"go-rest/internal/models"
	"go-rest/pkg/utils"
	"log"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	p = &models.Program{}
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#10b981")).Bold(true)

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
		errorMessage := errorStyle.Render("Error retrieving prohject name: " + err.Error())
		log.Fatalf(errorMessage)
	}


	err = promtDriverSelection()
	if err != nil {
		errorMessage := errorStyle.Render("Error retrieving drivers: " + err.Error())
		log.Fatalf(errorMessage)
	}

	baseDir := filepath.Join(".", p.Project)
	

	templData := models.TemplateData{
		Project: p.Project,
		Driver:  p.GetDriver(),
	}

	templates := config.PrepareTemplates(baseDir, p.GetDriver())

	for _, t := range templates {
		if err := utils.ValidateTemplate(t.TemplatePath); err != nil {
			errorMessage := errorStyle.Render(err.Error())
			panic(errorMessage)
		}
	}

	config.CreateProjectStructure(baseDir)
	fmt.Println()


	// Generate files in parallel
	statusChan := make(chan string)
	for _, tmplInfo := range templates {
		go utils.GenerateFileFromTemplate(tmplInfo.FilePath, tmplInfo.TemplatePath, templData, statusChan)
	}

	// Receive and print results
	for i := 0; i < len(templates); i++ {
		fmt.Println(<-statusChan)
	}

	fmt.Println()

	// Initialize Go module and tidy up
	if err := config.InitializeGoMod(p.Project); err != nil {
		log.Println(err)
	}

    // Create a success message
    successMessage := successStyle.Render("\nSuccessfully generated your Rest API: " + p.Project)
	fmt.Println(successMessage)
}
