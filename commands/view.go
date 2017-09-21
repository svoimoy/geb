package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/geb/commands/view"
)

// Tool:   geb
// Name:   view
// Usage:  view
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var ViewLong = `View information known to the geb tool.`

var ViewCmd = &cobra.Command{

	Use: "view",

	Aliases: []string{
		"v",
	},

	Short: "View information known to the geb tool.",

	Long: ViewLong,
}

func init() {
	RootCmd.AddCommand(ViewCmd)
}

func init() {
	// add sub-commands to this command when present

	ViewCmd.AddCommand(view.SystemCmd)
	ViewCmd.AddCommand(view.DslCmd)
	ViewCmd.AddCommand(view.GenCmd)
	ViewCmd.AddCommand(view.ProjectCmd)
	ViewCmd.AddCommand(view.DesignCmd)
	ViewCmd.AddCommand(view.PlansCmd)
}

// HOFSTADTER_BELOW
