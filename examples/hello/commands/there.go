package commands

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	"os"

	"github.com/spf13/cobra"
)

// Tool:   hello
// Name:   there
// Usage:
// Parent: hello

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var ThereCmd = &cobra.Command{

	Use: "there",

	Short: "say something to someone",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In thereCmd", "args", args)
		// Argument Parsing
		// [0]name:   who
		//     help:
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Printf("missing required argument: 'who'")
			cmd.Usage()
			os.Exit(1)
		}

		var who string

		if 0 < len(args) {

			who = args[0]
		}

		// [1]name:   what
		//     help:
		//     req'd:

		var what string

		if 1 < len(args) {

			what = args[1]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("hello there:",
			who,

			what,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(ThereCmd)
}

// HOFSTADTER_BELOW
