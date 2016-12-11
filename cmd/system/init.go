package system

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"

)

// Tool:   geb
// Name:   Init
// Usage:  init
// Parent: System
// ParentPath: 

var InitLong = `Intializes the geb tool and the ~/.hofstadter dot folder.`





var InitCmd = &cobra.Command {
	Use: "init",
	Aliases: []string{ 
		"initialize",
"setup",
	},
	Short: "Initialize the geb tool and files",
	Long: InitLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PersistentPreRun InitCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
	
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In InitCmd", "args", args)
		// Argument Parsing
		

		// HOFSTADTER_START cmd_run
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
