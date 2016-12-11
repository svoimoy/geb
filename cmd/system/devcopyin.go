package system

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"

)

// Tool:   geb
// Name:   DevCopyIn
// Usage:  dev-copy-in
// Parent: System
// ParentPath: 

var DevCopyInLong = `Copy development files to the dot folder`





var DevCopyInCmd = &cobra.Command {
	Use: "dev-copy-in",
	Short: "Copy development files to the dot folder",
	Long: DevCopyInLong,
		
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In DevCopyInCmd", "args", args)
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
long: Copy development files to the dot folder
name: DevCopyIn
parent: System
path: commands.subcommands
short: Copy development files to the dot folder
usage: dev-copy-in

*/
