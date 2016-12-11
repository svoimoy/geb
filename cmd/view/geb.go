package view

import (
	// HOFSTADTER_START import
	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"fmt"

	
	"github.com/spf13/cobra"

)

// Tool:   geb
// Name:   Geb
// Usage:  geb <dotpaths>...
// Parent: View
// ParentPath: 

var GebLong = `View information about the global geb config`





var GebCmd = &cobra.Command {
	Use: "geb <dotpaths>...",
	Short: "View information about Global geb config",
	Long: GebLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PersistentPreRun GebCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
	
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In GebCmd", "args", args)
		// Argument Parsing
		// [0]name:   paths
		//     help:   one ore more dotpaths for indexing into the data
		//     req'd:  
		var paths []string
			
		if 0 < len(args) {
			paths = args[0:]
		}
		
		

		// HOFSTADTER_START cmd_run
		ret, err := engine.ViewGeb(paths)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(ret)
		// HOFSTADTER_END   cmd_run
	},
		}


func init() {

}


/*
Repeated Context
----------------
args:
- help: one ore more dotpaths for indexing into the data
  name: paths
  rest: true
  type: array:string
long: View information about the global geb config
name: Geb
parent: View
path: commands.subcommands
short: View information about Global geb config
usage: geb <dotpaths>...

*/
