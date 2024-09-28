package utils

import (
	"fmt"
	"html/template"
	"os"
)

type TemplateInfo struct {
	FilePath     string
	TemplatePath string
}

// Generate file from a template
func GenerateFileFromTemplate(filePath, templatePath string, data interface{}, statusChan chan<- string) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		statusChan <- fmt.Sprintf("Failed to parse template %s: %v", templatePath, err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		statusChan <- fmt.Sprintf("Failed to create file %s: %v", filePath, err)
		return

	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		statusChan <- fmt.Sprintf("Failed to execute template %s: %v", templatePath, err)
		return
	}

	statusChan <- fmt.Sprintf("Successfully created file: %s", filePath)
}

// Validate that template file can be parsed successfully
func ValidateTemplate(templatePath string) error {
	_, err := template.ParseFiles(templatePath)
	return err
}