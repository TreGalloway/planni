package models

import (
	"time"
)

type Task struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	EstimatedTime float64   `json:"estimated_time"`
	ActualTime    float64   `json:"actual_time"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
}
