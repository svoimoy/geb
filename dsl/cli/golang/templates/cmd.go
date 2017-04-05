{{#with RepeatedContext as |RC| }}
{{#with DslContext as |CLI| }}
{{#if (eq RC.parent CLI.name) }}
package commands
{{else}}
package {{lower RC.parent}}
{{/if}}
// package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	// HOFSTADTER_START import
	{{#if RC.omit-run}}
	{{else}}
	"fmt"
	{{/if}}
	// HOFSTADTER_END   import

	{{#dotpath "args.required" RC false}}
	{{# with . as |D|}}
	{{#if (contains D "Error")}}
	{{else}}
	"os"
	{{/if}}
	{{/with}}
	{{/dotpath}}

	{{#if RC.flags}}
	"github.com/spf13/viper"
	{{else}}
		{{#if RC.pflags}}
	"github.com/spf13/viper"
		{{/if}}
	{{/if}}
	"github.com/spf13/cobra"

	{{#if subcommands}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/{{lower RC.name}}"
	{{/if}}
)

// Tool:   {{CLI.name}}
// Name:   {{RC.name}}
// Usage:  {{{RC.usage}}}
// Parent: {{{RC.parent}}}

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init


{{#if RC.long}}
var {{camelT RC.name}}Long = `{{{RC.long}}}`
{{/if}}

{{> "flag-var.go" RC }}

{{> "flag-init.go" RC }}

var {{camelT RC.name}}Cmd = &cobra.Command {
	{{#if RC.hidden}}
	Hidden: true,
	{{/if}}

	{{#if RC.usage}}
	Use: "{{{RC.usage}}}",
	{{else}}
	Use: "{{{RC.name}}}",
	{{/if}}

	{{#if RC.aliases}}
	Aliases: []string{ 
		{{#each RC.aliases}}"{{.}}",
		{{/each}}
	},
	{{/if}}

	{{#if RC.short}}
	Short: "{{{RC.short}}}",
	{{/if}}

	{{#if RC.long}}
	Long: {{camelT RC.name}}Long,
	{{/if}}

	{{#if RC.persistent-prerun}}
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PersistentPreRun {{RC.name}}Cmd", "args", args)
		{{> args-parse.go RC }}

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
	{{/if}}
	{{#if RC.prerun}}
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PreRun {{RC.name}}Cmd", "args", args)
		{{> args-parse.go RC }}

		// HOFSTADTER_START cmd_prerun
		// HOFSTADTER_END   cmd_prerun
	},
	{{/if}}
	{{#unless RC.omit-run}}
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In {{RC.name}}Cmd", "args", args)
		{{> args-parse.go RC }}

		// HOFSTADTER_START cmd_run
		fmt.Println("{{replace RC.pkgPath "/" " " -1}}:", {{#each RC.args}}
		{{camel name}},	
		{{/each}})
		// HOFSTADTER_END   cmd_run
	},
	{{/unless}}
	{{#if RC.persistent-postrun}}
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PersistentPostRun {{RC.name}}Cmd", "args", args)
		{{> args-parse.go RC }}

		// HOFSTADTER_START cmd_persistent_postrun
		// HOFSTADTER_END   cmd_persistent_postrun
	},
	{{/if}}
	{{#if RC.postrun}}
	PostRun: func(cmd *cobra.Command, args []string) {
		logger.Debug("In PostRun {{RC.name}}Cmd", "args", args)
		{{> args-parse.go RC }}

		// HOFSTADTER_START cmd_postrun
		// HOFSTADTER_END   cmd_postrun
	},
	{{/if}}
}


{{#if (eq RC.parent CLI.name) }}
func init() {
	RootCmd.AddCommand({{camelT RC.name}}Cmd)
}
{{/if}}

{{#if subcommands}}
func init() {
	{{#each subcommands}}
	{{camelT RC.name}}Cmd.AddCommand({{lower RC.name}}.{{camelT name}}Cmd)
	{{/each}}
}
{{/if}}

{{/with}}
{{/with}}

