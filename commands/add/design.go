package add

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   design
// Usage:  design [git-repo-url] <output-location>
// Parent: add

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var DesignCmd = &cobra.Command{

	Use: "design [git-repo-url] <output-location>",

	Short: "Add a design to a project.",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In designCmd", "args", args)
		// Argument Parsing
		// [0]name:   url
		//     help:   The url of a git repository. May be any of the remote types (git@, http(s)).
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'url'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var url string

		if 0 < len(args) {

			url = args[0]
		}

		// [1]name:   location
		//     help:   The location for the design. Defaults to the first design path listed in the geb.yaml file.
		//     req'd:

		var location string

		if 1 < len(args) {

			location = args[1]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb add design:",
			url,

			location,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}
