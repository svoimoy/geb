package templates

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   serveToolDB
// Name:   update
// Usage:  update <data-file>
// Parent: templates

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var UpdateCmd = &cobra.Command{

	Use: "update <data-file>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In updateCmd", "args", args)
		// Argument Parsing
		// [0]name:   data-file
		//     help:
		//     req'd:

		var dataFile string

		if 0 < len(args) {

			dataFile = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("serve-tool-db templates update:",
			dataFile,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
