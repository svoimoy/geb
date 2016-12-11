package view

import (
	// HOFSTADTER_START import
	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"fmt"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   System
// Usage:  sys <dotpaths>...
// Parent: View
// ParentPath:

var SystemLong = `View information about the global geb config`

var SystemCmd = &cobra.Command{
	Use: "sys <dotpaths>...",
	Aliases: []string{
		"s",
		"system",
		"geb",
		"config",
	},
	Short: "View information about Global geb config",
	Long:  SystemLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In SystemCmd", "args", args)
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
aliases:
- s
- system
- geb
- config
args:
- help: one ore more dotpaths for indexing into the data
  name: paths
  rest: true
  type: array:string
long: View information about the global geb config
name: System
parent: View
path: commands.subcommands
short: View information about Global geb config
usage: sys <dotpaths>...

*/
