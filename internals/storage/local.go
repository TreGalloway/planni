package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"tregalloway.com/planni/internal/models"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage() (*LocalStorage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	basePath := filepath.Join(homeDir, ".planni")
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, err
	}
	return &LocalStorage{
		basePath: basePath,
	}, nil
}
