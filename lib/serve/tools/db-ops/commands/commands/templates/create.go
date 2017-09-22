package templates

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports
	"github.com/hofstadter-io/geb/lib/types"

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   serveToolDB
// Name:   create
// Usage:  create <data-file>
// Parent: templates

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var CreateCmd = &cobra.Command{

	Use: "create <data-file>",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In createCmd", "args", args)
		// Argument Parsing
		// [0]name:   data-file
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'data-file'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var dataFile string

		if 0 < len(args) {

			dataFile = args[0]
		}

		var template types.Template
		// unmarshal data file into struct

	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
