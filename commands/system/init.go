package system

// package subcommands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   init
// Usage:  init
// Parent: system

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var InitLong = `Intializes the geb tool and the ~/.hofstadter dot folder.`

var InitCmd = &cobra.Command{
	Use: "init",
	Aliases: []string{
		"initialize",
		"setup",
	},
	Short: "Initialize the geb tool and files",
	Long:  InitLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In initCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		// HOFSTADTER_END   cmd_run
	},
}

func init() {

}

// HOFSTADTER_BELOW
