package utils

import (
	"fmt"
	"html/template"
	"os"
)

type fileGenerator struct {
	Path string
}

func NewFileGenerator(path string) *fileGenerator	{
	return &fileGenerator{
		Path: path,
	}
}

func (fg *fileGenerator) GenerateFolder() error {
	err := os.MkdirAll(fg.Path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create folder %s: %w", fg.Path, err)
	}
	fmt.Printf("Folder created at: %s\n", fg.Path)
	return nil
}

func (fg *fileGenerator) GenerateFile(templateData string) error {
	// Create the file
	file, err := os.Create(fg.Path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fg.Path, err)
	}
	defer file.Close()


	tmpl, err := template.New("file").Parse(templateData)
	if err != nil {
		return fmt.Errorf("error parsing template for %s: %v", fg.Path, err)
	}

	err = tmpl.Execute(file, nil) // You can pass data if needed
	if err != nil {
		return fmt.Errorf("failed to execute template for file %s: %w", fg.Path, err)
	}


	fmt.Printf("File created at: %s\n", fg.Path)
	return nil
}
