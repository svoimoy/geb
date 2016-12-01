{{#with RepeatedContext as |RC| }}
{{#with dsl.cli as |CLI| }}
package {{#lower RC.parent}}{{/lower}}

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	{{#if sub-commands}}
	"{{Proj.goimport_basedir}}/cmd/{{#lower RC.parent}}{{/lower}}/{{#lower RC.name}}{{/lower}}"
	{{/if}}

	{{#if RC.flags}}
	"github.com/spf13/viper"
	{{else}}
		{{#if RC.pflags}}
	"github.com/spf13/viper"
		{{/if}}
	{{/if}}
	"github.com/spf13/cobra"
)

// Tool:   {{CLI.name}}
// Name:   {{RC.name}}
// Usage:  {{RC.usage}}
// Parent: {{RC.parent}}

var {{RC.name}}Long = `{{long}}`

{{> "flag-var.go" RC }}

{{> "flag-init.go" RC }}

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
	{{#each sub-commands}}
	{{RC.name}}Cmd.AddCommand({{#lower RC.name}}{{/lower}}.{{name}}Cmd)
	{{/each}}
}

{{/with}}
{{/with}}
