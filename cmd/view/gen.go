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
// Name:   Gen
// Usage:  gen <dotpaths>...
// Parent: View
// ParentPath: 

var GenLong = `View information about generators known from the current path`





var GenCmd = &cobra.Command {
	Use: "gen <dotpaths>...",
	Short: "View information about Generators",
	Long: GenLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PersistentPreRun GenCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
	
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In GenCmd", "args", args)
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

		ret, err := engine.ViewGen(cwd, paths)
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
long: View information about generators known from the current path
name: Gen
parent: View
path: commands.subcommands
short: View information about Generators
usage: gen <dotpaths>...

*/
