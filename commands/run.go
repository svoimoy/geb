package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   run
// Usage:  run <run-onfigname>
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var RunLong = `Run a run-config pipeline specified in your project  file.
Use this to run pre and post steps around 'gen gen'.
run-config Pipelines may also be used by generators.
See [docs link...] for more information.
`

var RunCmd = &cobra.Command{

	Use: "run <run-onfigname>",

	Aliases: []string{
		"b",
	},

	Short: "Run a run-config pipeline for a project.",

	Long: RunLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In runCmd", "args", args)
		// Argument Parsing
		// [0]name:   stages
		//     help:   The stages to run in order. Used to override the pipeline in the project file.
		//     req'd:

		var stages []string

		if 0 < len(args) {
			stages = args[0:]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb run:",
			stages,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
