{{#with RepeatedContext as |RC| }}
{{#with dsl.cli as |CLI| }}
package cmd

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   {{CLI.name}}
// Name:   {{RC.name}}
// Usage:  {{RC.usage}}
// Parent: {{RC.parent}}

var {{RC.name}}Long = `{{long}}`

var {{RC.name}}Cmd = &cobra.Command {
	Use: "{{RC.usage}}",
	Short: "{{RC.short}}",
	Long: {{RC.name}}Long,
	Run: func(cmd *cobra.Command, args []string) {
		// HOFSTADTER_START cmd_run
		fmt.Println("In {{RC.name}}Cmd")
		// HOFSTADTER_END   cmd_run
	},
}


func init() {
	{{RC.parent}}Cmd.AddCommand({{RC.name}}Cmd)
}

{{/with}}
{{/with}}
