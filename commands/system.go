package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/geb/commands/system"
)

// Tool:   geb
// Name:   system
// Usage:  system
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var SystemLong = `Manage the geb system and congiuration`

var SystemCmd = &cobra.Command{

	Use: "system",

	Aliases: []string{
		"sys",
		"s",
	},

	Short: "Manage the geb system and congiuration",

	Long: SystemLong,
}

func init() {
	RootCmd.AddCommand(SystemCmd)
}

func init() {
	// add sub-commands to this command when present

	SystemCmd.AddCommand(system.InitCmd)
	SystemCmd.AddCommand(system.DevModeCmd)
	SystemCmd.AddCommand(system.UpdateCmd)
}

// HOFSTADTER_BELOW
