{{#with RepeatedContext as |RC| }}
{{#with dsl.cli as |CLI| }}
package cmd

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	{{#if subcommands}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/{{lower RC.name}}"
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
// Usage:  {{{RC.usage}}}
// Parent: {{{RC.parent}}}

var {{RC.name}}Long = `{{{RC.long}}}`

{{> "flag-var.go" RC }}

{{> "flag-init.go" RC }}

var {{RC.name}}Cmd = &cobra.Command {
	Use: "{{{RC.usage}}}",
	{{#if RC.aliases}}
	Aliases: []string{ 
		{{#each RC.aliases}}"{{.}}",
		{{/each}}
	},
	{{/if}}
	Short: "{{{RC.short}}}",
	Long: {{RC.name}}Long,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("In {{RC.name}}Cmd", args)
		{{> args-parse.go RC }}

		// HOFSTADTER_START cmd_run
		// HOFSTADTER_END   cmd_run
	},
}


func init() {
	{{#if (eq RC.parent CLI.name) }}
	RootCmd.AddCommand({{RC.name}}Cmd)
	{{else}}
	{{RC.parent}}Cmd.AddCommand({{RC.name}}Cmd)
	{{/if}}

	{{#each subcommands}}
	{{RC.name}}Cmd.AddCommand({{lower RC.name}}.{{name}}Cmd)
	{{/each}}
}

{{/with}}
{{/with}}

/*
Repeated Context
----------------
{{{yaml RepeatedContext}}}
*/
