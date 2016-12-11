package cmd

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Init
// Usage:  init
// Parent: geb

var InitLong = `Intializes the geb tool and the ~/.hofstadter dotfolder`





var InitCmd = &cobra.Command {
	Use: "init",
	Aliases: []string{ 
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
		fmt.Println("In InitCmd")
		// HOFSTADTER_END   cmd_run
	},
		}


func init() {
	RootCmd.AddCommand(InitCmd)


}


/*
Repeated Context
----------------
aliases:
- setup
long: Intializes the geb tool and the ~/.hofstadter dotfolder
name: Init
parent: geb
path: commands
short: Initialize the geb tool and files
usage: init

*/
