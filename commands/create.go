package commands

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	"os"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   create
// Usage:  create <name> <dsl/gen>...
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var CreateLong = `Create a new project with the given name.
Optionally specifiy the starting set of
DSLs and generators for the project.
The output directory defaults to the same name,
unless overridden by the output flag.
`

var CreateCmd = &cobra.Command{

	Use: "create <name> <dsl/gen>...",

	Aliases: []string{
		"new",
	},

	Short: "Create a new project",

	Long: CreateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In createCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:   The name of the new project to create
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Printf("missing required argument: 'name'")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		// [1]name:   dsls_n_gens
		//     help:   The starting list of DSLs and generators by path.
		//     req'd:

		var dslsNGens []string

		if 1 < len(args) {
			dslsNGens = args[1:]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb create:", name, dslsNGens)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(CreateCmd)
}

// HOFSTADTER_BELOW
