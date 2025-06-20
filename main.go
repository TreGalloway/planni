package main

import (
	"context"
	"os"
	"github.com/charmbracelet/fang"	
	"github.com/TreGalloway/planni/cmd"
)

func main() {
	if err := fang.Execute(context.Background(), cmd.GetRoot()); err != nil {
		os.Exit(1)
	}
}
