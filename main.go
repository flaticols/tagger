package main

import (
	"github.com/flaticols/tagger/commands"
	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	root.AddCommand(commands.DoCommand())

	_, err := root.ExecuteC()
	if err != nil {
		return
	}
}
