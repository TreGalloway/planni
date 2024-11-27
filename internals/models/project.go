package models

import (
	"time"
)

type Project struct {
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Tasks   []Task    `json:"tasks"`
}
