package ui

import "github.com/charmbracelet/lipgloss"

var (
	StatusStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1)

	TaskStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true)
)
