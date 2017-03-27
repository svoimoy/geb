package commands

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   build
// Usage:  build
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var BuildLong = `Run the build pipeline specified in your project.
Use this to run pre and post steps around 'gen gen'.
Pipelines are also used by generators.
See [...] for more information.
`

var BuildCmd = &cobra.Command{
	Use: "build",
	Aliases: []string{
		"b",
	},
	Short: "Run the build pipeline for a project.",
	Long:  BuildLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In buildCmd", "args", args)
		// Argument Parsing
		// [0]name:   stages
		//     help:   The stages to run in order. Used to override the pipeline in the project file.
		//     req'd:
		var stages []string

		if 0 < len(args) {
			stages = args[0:]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb build:", stages)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(BuildCmd)

}

// HOFSTADTER_BELOW
