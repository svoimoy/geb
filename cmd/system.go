package cmd

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/cmd/system"
)

// Tool:   geb
// Name:   System
// Usage:  system
// Parent: geb

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

	system.SetLogger(logger)
	SystemCmd.AddCommand(system.InitCmd)
	SystemCmd.AddCommand(system.DevCopyInCmd)
	SystemCmd.AddCommand(system.UpdateCmd)

}

// HOFSTADTER_BELOW
