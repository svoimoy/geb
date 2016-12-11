package view

import (
	// HOFSTADTER_START import
	"github.ibm.com/hofstadter-io/geb/engine"
	"os"
	// HOFSTADTER_END   import

	"fmt"

	
	"github.com/spf13/cobra"

)

// Tool:   geb
// Name:   Dsl
// Usage:  dsl <dotpath>...
// Parent: View
// ParentPath: 

var DslLong = `View information about DSLs known from the current path`





var DslCmd = &cobra.Command {
	Use: "dsl <dotpath>...",
	Short: "View information about DSLs",
	Long: DslLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PersistentPreRun DslCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
	
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In DslCmd", "args", args)
		// Argument Parsing
		// [0]name:   paths
		//     help:   one ore more dotpaths for indexing into the data
		//     req'd:  
		var paths []string
			
		if 0 < len(args) {
			paths = args[0:]
		}
		
		

		// HOFSTADTER_START cmd_run
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		ret, err := engine.ViewDsl(cwd, paths)
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
long: View information about DSLs known from the current path
name: Dsl
parent: View
path: commands.subcommands
short: View information about DSLs
usage: dsl <dotpath>...

*/
