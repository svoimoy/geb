package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/geb/commands/gebberish"
)

// Tool:   geb
// Name:   gebberish
// Usage:  gebberish
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var GebberishLong = `Games, shenanigans, and other gebberish.`

var GebberishCmd = &cobra.Command{
	Hidden: true,

	Use: "gebberish",

	Aliases: []string{
		"games",
		"G",
	},

	Short: "it's a puzzle?!",

	Long: GebberishLong,
}

func init() {
	RootCmd.AddCommand(GebberishCmd)
}

func init() {
	// add sub-commands to this command when present

	GebberishCmd.AddCommand(gebberish.MiCmd)
}

// HOFSTADTER_BELOW
