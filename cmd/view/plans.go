package view

import (
	// HOFSTADTER_START import
	"fmt"
	"os"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Plans
// Usage:  plans <dotpath>...
// Parent: View
// ParentPath:

var PlansLong = `View information about a Project's Plans known from the current path`

var PlansCmd = &cobra.Command{
	Use: "plans <dotpath>...",
	Aliases: []string{
		"P",
	},
	Short: "View information about a Project's Plans",
	Long:  PlansLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PlansCmd", "args", args)
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

// HOFSTADTER_BELOW
