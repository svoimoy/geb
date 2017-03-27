package commands

// package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/commands/system"
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
	Long:  SystemLong,
}

func init() {
	RootCmd.AddCommand(SystemCmd)

	SystemCmd.AddCommand(system.InitCmd)
	SystemCmd.AddCommand(system.DevCopyInCmd)
	SystemCmd.AddCommand(system.UpdateCmd)

}

// HOFSTADTER_BELOW
