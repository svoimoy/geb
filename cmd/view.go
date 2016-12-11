package cmd

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/cmd/view"
)

// Tool:   geb
// Name:   View
// Usage:  view
// Parent: geb

var ViewLong = `View information known to the geb tool.`





var ViewCmd = &cobra.Command {
	Use: "view",
	Aliases: []string{ 
		"v",
	},
	Short: "View information known to the geb tool.",
	Long: ViewLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		view.SetLogger(logger)
		logger.Debug("In PersistentPreRun ViewCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
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

	view.SetLogger(logger)
}


/*
Repeated Context
----------------
aliases:
- v
long: View information known to the geb tool.
name: View
omit-run: true
parent: geb
path: commands
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
