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

var ViewCmd = &cobra.Command{
	Use: "view",
	Aliases: []string{
		"v",
	},
	Short: "View information known to the geb tool.",
	Long:  ViewLong,
}

func init() {
	RootCmd.AddCommand(ViewCmd)

	ViewCmd.AddCommand(view.SystemCmd)
	ViewCmd.AddCommand(view.DslCmd)
	ViewCmd.AddCommand(view.GenCmd)
	ViewCmd.AddCommand(view.ProjectCmd)
	ViewCmd.AddCommand(view.DesignCmd)
	ViewCmd.AddCommand(view.PlansCmd)
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
- aliases:
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
- aliases:
  - d
  args:
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
- aliases:
  - g
  args:
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
  - p
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
- aliases:
  - D
  args:
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
- aliases:
  - P
  args:
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
