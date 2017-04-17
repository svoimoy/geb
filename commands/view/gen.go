package view

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	"os"
	"path/filepath"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   gen
// Usage:  gen <dotpaths>...
// Parent: view

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var GenLong = `View information about generators known from the current path`

var GenCmd = &cobra.Command{

	Use: "gen <dotpaths>...",

	Aliases: []string{
		"g",
	},

	Short: "View information about Generators",

	Long: GenLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In genCmd", "args", args)
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
		d_dir := filepath.Join(cwd, "design")
		ret, err := engine.ViewGen(d_dir, paths)
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
