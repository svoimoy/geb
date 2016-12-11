package system

import (
	// HOFSTADTER_START import
	"github.ibm.com/hofstadter-io/geb/engine/system"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Init
// Usage:  init
// Parent: System
// ParentPath:

var InitLong = `Intializes the geb tool and the ~/.hofstadter dot folder.`

var InitCmd = &cobra.Command{
	Use: "init",
	Aliases: []string{
		"initialize",
		"setup",
	},
	Short: "Initialize the geb tool and files",
	Long:  InitLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In InitCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		system.Init()
		// HOFSTADTER_END   cmd_run
	},
}

func init() {

}

/*
Repeated Context
----------------
aliases:
- initialize
- setup
long: Intializes the geb tool and the ~/.hofstadter dot folder.
name: Init
parent: System
path: commands.subcommands
short: Initialize the geb tool and files
usage: init

*/
