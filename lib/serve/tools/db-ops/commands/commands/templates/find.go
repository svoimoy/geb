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
// Name:   find
// Usage:  find <uuid>
// Parent: templates

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var FindCmd = &cobra.Command{

	Use: "find <uuid>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In findCmd", "args", args)
		// Argument Parsing
		// [0]name:   uuid
		//     help:
		//     req'd:

		var uuid string

		if 0 < len(args) {

			uuid = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("serve-tool-db templates find:",
			id,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
