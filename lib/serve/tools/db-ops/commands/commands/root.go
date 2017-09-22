package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	// infered imports

	"github.com/spf13/cobra"
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var (
	RootCmd = &cobra.Command{

		Use: "serveToolDB",

		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("In serveToolDBCmd", "args", args)
			// Argument Parsing

			// HOFSTADTER_START cmd_run
			// HOFSTADTER_END   cmd_run
		},
	}
)

// HOFSTADTER_BELOW
