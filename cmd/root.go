// planni/cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "planni",
	Short: "Your terminal backbone for productivity.",
	Long: `Planni is a modern, all-in-one CLI designed to enhance your terminal
workflow with smart navigation, rich file listings, and powerful utilities.`,
	// This function will run if no subcommands are specified
	Run: func(cmd *cobra.Command, args []string) {
		// For now, we can just show the help screen.
		// Later, this could default to the 'ls' command.
		cmd.Help()
	},
}
func GetRoot() *cobra.Command {
	return rootCmd
}
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
