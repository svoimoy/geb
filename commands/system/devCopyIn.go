package system

// package subcommands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   dev-copy-in
// Usage:  dev-copy-in
// Parent: system

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var DevCopyInLong = `Copy development files to the dot folder`

var DevCopyInCmd = &cobra.Command{

	Use: "dev-copy-in",

	Short: "Copy development files to the dot folder",

	Long: DevCopyInLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In dev-copy-inCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		// HOFSTADTER_END   cmd_run
	},
}

// HOFSTADTER_BELOW
