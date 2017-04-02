package commands

// package commands

import (
	// HOFSTADTER_START import
	"fmt"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/commands/gen"
)

// Tool:   geb
// Name:   gen
// Usage:  gen
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var GenLong = `Generate a project from its working directory.`

var GenCmd = &cobra.Command{
	Use: "gen",
	Aliases: []string{
		"geb",
		"geberate",
		"generate",
		"g",
	},
	Short: "Generate a project.",
	Long:  GenLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In genCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		filename := "geb.yaml"

		err := engine.GenerateProject(filename, args)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(GenCmd)

	GenCmd.AddCommand(gen.FileCmd)

}

// HOFSTADTER_BELOW