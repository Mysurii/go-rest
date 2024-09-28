package config

import (
	"fmt"
	"go-rest/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateProjectStructure(baseDir string) {
	dirs := []string{
		filepath.Join(baseDir, "cmd", "api"),
		filepath.Join(baseDir, "internal", "server"),
		filepath.Join(baseDir, "internal", "database"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}
}

func PrepareTemplates(baseDir, driver string) []utils.TemplateInfo {
	return []utils.TemplateInfo{
		{FilePath: filepath.Join(baseDir, "cmd", "api", "main.go"), TemplatePath: "internal/templates/cmd.api.main.go.templ"},
		{FilePath: filepath.Join(baseDir, "internal", "server", "server.go"), TemplatePath: "internal/templates/server.server.go.templ"},
		{FilePath: filepath.Join(baseDir, "internal", "server", "routes.go"), TemplatePath: "internal/templates/server.routes.go.templ"},
		{FilePath: filepath.Join(baseDir, "internal", "database", "database.go"), TemplatePath: fmt.Sprintf("internal/templates/db.%s.go.templ", strings.ToLower(driver))},
		{FilePath: filepath.Join(baseDir, ".env"), TemplatePath: "internal/templates/env.templ"},
	}
}

func InitializeGoMod(projectName string) error {
	if err := utils.RunGoModInit(projectName); err != nil {
		return err
	}

	return utils.RunGoModTidy(projectName)
}