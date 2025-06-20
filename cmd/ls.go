// in cmd/ls.go
package cmd

import (
	"fmt"
	"log" // We'll use the log package for fatal errors
	"os"  // We need the os package to interact with the filesystem

	"github.com/spf13/cobra"
	"github.com/TreGalloway/planni/internal/tui"
)

var lsCmd = &cobra.Command{
	Use:   "ls [path]",
	Short: "A modern, beautiful file listing (eza-like)",
	Long:  `Provides a rich file listing with icons, git status, and other modern features.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Get our default styles.
		styles := tui.DefaultStyles()

		// 2. Determine the target directory.
		var targetDir string
		if len(args) > 0 {
			targetDir = args[0]
		} else {
			// If no path is provided, use the current working directory.
			var err error
			targetDir, err = os.Getwd()
			if err != nil {
				log.Fatalf("Error getting current directory: %v", err)
			}
		}

		// 3. Read the directory contents.
		entries, err := os.ReadDir(targetDir)
		if err != nil {
			log.Fatalf("Error reading directory %s: %v", targetDir, err)
		}

		// 4. Loop through each entry and print it with the correct style.
		for _, entry := range entries {
			if entry.IsDir() {
				// It's a directory, apply the directory style.
				// We'll add a trailing slash for clarity, like `ls -F`.
				name := entry.Name() + "/"
				fmt.Println(styles.Directory.Render(name))
			} else {
				// It's a file, apply the file style.
				name := entry.Name()
				fmt.Println(styles.File.Render(name))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

