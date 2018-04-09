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
// Usage:  run <run-config-pipeline-name>...
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var RunLong = `Run a run-config pipeline specified in your project  file.
Use this to run pre and post steps around 'gen gen'.
See [docs link...] for more information.
`

var RunCmd = &cobra.Command{

	Use: "run <run-config-pipeline-name>...",

	Short: "Run a run-config pipeline for a project.",

	Long: RunLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In runCmd", "args", args)
		// Argument Parsing
		// [0]name:   pipelines
		//     help:   The pipelines to run, in order.
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'pipelines'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var pipelines []string

		if 0 < len(args) {
			pipelines = args[0:]
		}

		// HOFSTADTER_START cmd_run
		err := run.Run(pipelines)
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
