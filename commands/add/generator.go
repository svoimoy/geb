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
// Name:   generator
// Usage:  generator [git-repo-url] <output-location>
// Parent: add

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var GeneratorCmd = &cobra.Command{

	Use: "generator [git-repo-url] <output-location>",

	Aliases: []string{
		"gen",
	},

	Short: "Add a generator to a project.",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In generatorCmd", "args", args)
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
		//     help:   The location for the generator. Defaults to the first dsl path listed in the geb.yaml file.
		//     req'd:

		var location string

		if 1 < len(args) {

			location = args[1]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb add generator:",
			url,

			location,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}
