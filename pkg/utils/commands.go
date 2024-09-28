package utils

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunGoModInit(projectName string) error {
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = filepath.Join(".", projectName) // Set the directory to the project folder

	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "already exists") {
			fmt.Printf("Module already exists for %s. Skipping go mod init.\n", projectName)
			return nil // Return nil if the module already exists
		}
		return fmt.Errorf("go mod init failed: %v, output: %s", err, string(output))
	}

	fmt.Println(string(output))
	return nil
}

func RunGoModTidy(projectName string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = filepath.Join(".", projectName) // Set the directory to the project folder

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod tidy failed: %v, output: %s", err, string(output))
	}

	fmt.Println(string(output))
	return nil
}
