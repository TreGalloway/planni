package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "planni",
	Short: "Planni - A CLI for planning sprints and managing tasks",
	Long: `Planni helps you manage your projects with an intuitive CLI interface.
It supports sprint planning, task management, and visual roadmaps.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags can be added here
}
