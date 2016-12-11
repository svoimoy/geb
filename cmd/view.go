package cmd

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	"github.ibm.com/hofstadter-io/geb/cmd/view"

	
	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   View
// Usage:  view
// Parent: geb

var ViewLong = `View information known to the geb tool.`





var ViewCmd = &cobra.Command {
	Use: "view",
	Short: "View information known to the geb tool.",
	Long: ViewLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("In ViewCmd", args)
		// Argument Parsing
		

		// HOFSTADTER_START cmd_run
		fmt.Println("In ViewCmd")
		// HOFSTADTER_END   cmd_run
	},
}


func init() {
	RootCmd.AddCommand(ViewCmd)

	ViewCmd.AddCommand(view.GebCmd)
	ViewCmd.AddCommand(view.DslCmd)
	ViewCmd.AddCommand(view.GenCmd)
	ViewCmd.AddCommand(view.ProjectCmd)
	ViewCmd.AddCommand(view.DesignCmd)
	ViewCmd.AddCommand(view.PlansCmd)
}


/*
Repeated Context
----------------
long: View information known to the geb tool.
name: View
parent: geb
path: commands
run: false
short: View information known to the geb tool.
subcommands:
- args:
  - help: one ore more dotpaths for indexing into the data
    name: paths
    rest: true
    type: array:string
  long: View information about the global geb config
  name: Geb
  parent: View
  path: commands.subcommands
  short: View information about Global geb config
  usage: geb <dotpaths>...
- args:
  - help: one ore more dotpaths for indexing into the data
    name: paths
    rest: true
    type: array:string
  long: View information about DSLs known from the current path
  name: Dsl
  parent: View
  path: commands.subcommands
  short: View information about DSLs
  usage: dsl <dotpath>...
- args:
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
- aliases:
  - proj
  args:
  - help: one ore more dotpaths for indexing into the data
    name: paths
    rest: true
    type: array:string
  long: View information about a Project known from the current path
  name: Project
  parent: View
  path: commands.subcommands
  short: View information about a Project
  usage: project <dotpath>...
- args:
  - help: one ore more dotpaths for indexing into the data
    name: paths
    rest: true
    type: array:string
  long: View information about Designs known from the current path
  name: Design
  parent: View
  path: commands.subcommands
  short: View information about Designs
  usage: design <dotpath>...
- args:
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
usage: view

*/
