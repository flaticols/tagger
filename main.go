package main

import (
	"github.com/flaticols/tagger/commands"
	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	root.AddCommand(commands.CreateCommand())

	_, err := root.ExecuteC()
	if err != nil {
		return
	}
}
