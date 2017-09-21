package view

import (
	// HOFSTADTER_START import
	"fmt"

	"github.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   system
// Usage:  sys <dotpaths>...
// Parent: view

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

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

	Long: SystemLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In systemCmd", "args", args)
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
		fmt.Println(ret)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
