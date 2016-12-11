package view

import (
	// HOFSTADTER_START import
	"fmt"
	"github.ibm.com/hofstadter-io/geb/engine"
	"os"
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Plans
// Usage:  plans <dotpath>...
// Parent: View
// ParentPath: 

var PlansLong = `View information about a Project's Plans known from the current path`





var PlansCmd = &cobra.Command {
	Use: "plans <dotpath>...",
	Short: "View information about a Project's Plans",
	Long: PlansLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("In PlansCmd", args)
		// Argument Parsing
		// [0]name:   paths
		//     help:   one ore more dotpaths for indexing into the data
		//     req'd:  
		var paths []string
			
		if 0 < len(args) {
			paths = args[0:]
		}
		
		fmt.Println("arg[0] = ", paths)
		
		

		// HOFSTADTER_START cmd_run
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		ret, err := engine.ViewPlans(cwd, paths)
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
long: View information about a Project's Plans known from the current path
name: Plans
parent: View
path: commands.subcommands
short: View information about a Project's Plans
usage: plans <dotpath>...

*/
