package cmd

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/cmd/gebberish"
)

// Tool:   geb
// Name:   Gebberish
// Usage:  gebberish
// Parent: geb

var GebberishLong = `Games, shenanigans, and other gebberish.`

var GebberishCmd = &cobra.Command{
	Hidden: true,
	Use:    "gebberish",
	Aliases: []string{
		"games",
		"G",
	},
	Short: "it's a puzzle?!",
	Long:  GebberishLong,
}

func init() {
	RootCmd.AddCommand(GebberishCmd)

	gebberish.SetLogger(logger)
	GebberishCmd.AddCommand(gebberish.MiCmd)

}

// HOFSTADTER_BELOW
