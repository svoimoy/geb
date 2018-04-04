package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/geb/commands/new"
)

// Tool:   geb
// Name:   new
// Usage:  new
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var NewLong = `Create new geb projects, dsls, generators, and designs.`

var NewCmd = &cobra.Command{

	Use: "new",

	Short: "Create new stuff.",

	Long: NewLong,
}

func init() {
	RootCmd.AddCommand(NewCmd)
}

func init() {
	// add sub-commands to this command when present

	NewCmd.AddCommand(new.ProjectCmd)
	NewCmd.AddCommand(new.GeneratorCmd)
	NewCmd.AddCommand(new.DslCmd)
}

// HOFSTADTER_BELOW
