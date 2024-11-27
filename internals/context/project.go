package context

import (
	"fmt"
	"os"
	"path/filepath"
)

type ProjectContext struct {
	Path       string
	ConfigPath string
}

func FindProjectContext() (*ProjectContext, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath, err := findConfigFile(currentDir)
	if err != nil {
		return nil, err
	}

	return &ProjectContext{
		Path:       filepath.Dir(configPath),
		ConfigPath: configPath,
	}, nil
}

func findConfigFile(startPath string) (string, error) {
	currentPath := startPath
	for {
		configPath := filepath.Join(currentPath, ".planni.json")
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}

		parentPath := filepath.Dir(currentPath)
		if parentPath == currentPath {
			return "", fmt.Errorf("no .planni.json found in current directory or any parent")
		}
		currentPath = parentPath
	}
}
