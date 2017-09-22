package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/geb/lib/serve/tools/db-ops/commands/commands/templates"
)

// Tool:   serveToolDB
// Name:   templates
// Usage:
// Parent: serveToolDB

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var TemplatesCmd = &cobra.Command{

	Use: "templates",

	Short: "work with the templates resource",
}

func init() {
	RootCmd.AddCommand(TemplatesCmd)
}

func init() {
	// add sub-commands to this command when present

	TemplatesCmd.AddCommand(templates.MigrateCmd)
	TemplatesCmd.AddCommand(templates.CreateCmd)
	TemplatesCmd.AddCommand(templates.FindCmd)
	TemplatesCmd.AddCommand(templates.UpdateCmd)
	TemplatesCmd.AddCommand(templates.DeleteCmd)
	TemplatesCmd.AddCommand(templates.QueryCmd)
}

// HOFSTADTER_BELOW
