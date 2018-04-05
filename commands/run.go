package commands

import (
	// HOFSTADTER_START import
	"fmt"

	"github.com/hofstadter-io/geb/lib/run"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	"os"

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

	Short: "Run a run-config pipeline for a project.",

	Long: RunLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In runCmd", "args", args)
		// Argument Parsing
		// [0]name:   configs
		//     help:   The stages to run in order. Used to override the pipeline in the project file.
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'configs'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var configs []string

		if 0 < len(args) {
			configs = args[0:]
		}

		// HOFSTADTER_START cmd_run
		err := run.Run(configs)
		if err != nil {
			fmt.Println("Error:", err)
		}
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
