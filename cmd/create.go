package cmd

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Create
// Usage:  create <name> <dsl/gen>...
// Parent: geb

var CreateLong = `Create a new project with the given name.
Optionally specifiy the starting set of
DSLs and generators for the project.
The output directory defaults to the same name,
unless overridden by the output flag.
`





var CreateCmd = &cobra.Command {
	Use: "create <name> <dsl/gen>...",
	Short: "Create a new project",
	Long: CreateLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("In CreateCmd", args)
		// Argument Parsing
		// [0]name:   name
		//     help:   The name of the new project to create
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("not enough args supplied")
			return
		}
		var name string
		if 0 < len(args) {
			name = args[0]
		}
		
		fmt.Println("arg[0] = ", name)
		
		// [1]name:   dsls_n_gens
		//     help:   The starting list of DSLs and generators by path.
		//     req'd:  
		var dsls_n_gens []string
			
		if 1 < len(args) {
			dsls_n_gens = args[1:]
		}
		
		fmt.Println("arg[1] = ", dsls_n_gens)
		
		

		// HOFSTADTER_START cmd_run
		fmt.Println("In CreateCmd")
		// HOFSTADTER_END   cmd_run
	},
}


func init() {
	RootCmd.AddCommand(CreateCmd)

}


/*
Repeated Context
----------------
args:
- help: The name of the new project to create
  name: name
  required: true
  type: string
- help: The starting list of DSLs and generators by path.
  name: dsls_n_gens
  rest: true
  type: array:string
long: |
  Create a new project with the given name.
  Optionally specifiy the starting set of
  DSLs and generators for the project.
  The output directory defaults to the same name,
  unless overridden by the output flag.
name: Create
parent: geb
path: commands
short: Create a new project
usage: create <name> <dsl/gen>...

*/
