package commands

import (
	// HOFSTADTER_START import
	"fmt"

	"github.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
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

	Use: "gen [subdesign-dotpaths]",

	Aliases: []string{
		"geb",
		"geberate",
		"generate",
		"g",
	},

	Short: "Generate a project.",

	Long: GenLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In genCmd", "args", args)
		// Argument Parsing
		// [0]name:   paths
		//     help:   A list of filepaths to use as the designs. If no paths are provided, than the geb.yaml is consulted.
		//     req'd:

		var paths []string

		if 0 < len(args) {
			paths = args[0:]
		}

		// HOFSTADTER_START cmd_run
		filename := "geb.yaml"

		err := engine.GenerateProject(filename, paths)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(GenCmd)
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
