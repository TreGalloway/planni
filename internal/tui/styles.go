package tui

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Directory lipgloss.Style
	File	lipgloss.Style
}

func DefaultStyles() *Styles {
	dirStyle := lipgloss.NewStyle().
	Bold(true).
		Foreground(lipgloss.AdaptiveColor{Light: "#277BC0", Dark: "#78A2CC"})

	fileStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#434242", Dark: "#C8C6C6"})

return &Styles{
		Directory: dirStyle,
		File: fileStyle,
}
}
