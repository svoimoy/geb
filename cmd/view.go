package cmd

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/cmd/view"
)

// Tool:   geb
// Name:   View
// Usage:  view
// Parent: geb

var ViewLong = `View information known to the geb tool.`

var ViewCmd = &cobra.Command{
	Use: "view",
	Aliases: []string{
		"v",
	},
	Short: "View information known to the geb tool.",
	Long:  ViewLong,
}

func init() {
	RootCmd.AddCommand(ViewCmd)

	ViewCmd.AddCommand(view.SystemCmd)
	ViewCmd.AddCommand(view.DslCmd)
	ViewCmd.AddCommand(view.GenCmd)
	ViewCmd.AddCommand(view.ProjectCmd)
	ViewCmd.AddCommand(view.DesignCmd)
	ViewCmd.AddCommand(view.PlansCmd)

}

// HOFSTADTER_BELOW
