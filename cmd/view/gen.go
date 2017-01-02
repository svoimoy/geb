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

var GenCmd = &cobra.Command{
	Use: "gen <dotpaths>...",
	Aliases: []string{
		"g",
	},
	Short: "View information about Generators",
	Long:  GenLong,

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

// HOFSTADTER_BELOW
