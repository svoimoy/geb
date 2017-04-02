package system

// package subcommands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   update
// Usage:  update
// Parent: system

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var UpdateLong = `Update the geb library DSLs, designs, and other files in the dot folder.`

var UpdateCmd = &cobra.Command{

	Use: "update",

	Short: "Update the geb library and dot folder",

	Long: UpdateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In updateCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		// HOFSTADTER_END   cmd_run
	},
}

// HOFSTADTER_BELOW
