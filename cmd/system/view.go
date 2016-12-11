package system

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   View
// Usage:  view
// Parent: System
// ParentPath:

var ViewLong = `View the geb system configuration`

var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the geb system configuration",
	Long:  ViewLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In ViewCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		// HOFSTADTER_END   cmd_run
	},
}

func init() {

}

/*
Repeated Context
----------------
long: View the geb system configuration
name: View
parent: System
path: commands.subcommands
short: View the geb system configuration
usage: view

*/
