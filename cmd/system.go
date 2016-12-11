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





var SystemCmd = &cobra.Command {
	Use: "system",
	Aliases: []string{ 
		"sys",
	},
	Short: "Manage the geb system and congiuration",
	Long: SystemLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		system.SetLogger(logger)
		logger.Debug("In PersistentPreRun SystemCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
			}


func init() {
	RootCmd.AddCommand(SystemCmd)

	SystemCmd.AddCommand(system.InitCmd)
	SystemCmd.AddCommand(system.DevCopyInCmd)
	SystemCmd.AddCommand(system.UpdateCmd)
	SystemCmd.AddCommand(system.ViewCmd)

	system.SetLogger(logger)
}


/*
Repeated Context
----------------
aliases:
- sys
long: Manage the geb system and congiuration
name: System
omit-run: true
parent: geb
path: commands
short: Manage the geb system and congiuration
subcommands:
- aliases:
  - initialize
  - setup
  long: Intializes the geb tool and the ~/.hofstadter dot folder.
  name: Init
  parent: System
  path: commands.subcommands
  short: Initialize the geb tool and files
  usage: init
- long: Copy development files to the dot folder
  name: DevCopyIn
  parent: System
  path: commands.subcommands
  short: Copy development files to the dot folder
  usage: dev-copy-in
- long: Update the geb library DSLs, designs, and other files in the dot folder.
  name: Update
  parent: System
  path: commands.subcommands
  short: Update the geb library and dot folder
  usage: update
- long: View the geb system configuration
  name: View
  parent: System
  path: commands.subcommands
  short: View the geb system configuration
  usage: view
usage: system

*/
