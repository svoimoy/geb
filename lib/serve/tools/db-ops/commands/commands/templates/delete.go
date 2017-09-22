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
// Name:   delete
// Usage:  delete <uuid>
// Parent: templates

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var DeleteCmd = &cobra.Command{

	Use: "delete <uuid>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In deleteCmd", "args", args)
		// Argument Parsing
		// [0]name:   uuid
		//     help:
		//     req'd:

		var uuid string

		if 0 < len(args) {

			uuid = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("serve-tool-db templates delete:",
			id,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
