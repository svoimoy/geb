{{#with DslContext as |CLI| }}
package commands

import (
  // HOFSTADTER_START import
  // HOFSTADTER_END   import

	{{#if CLI.flags}}
	"github.com/spf13/viper"
	{{else}}
		{{#if CLI.pflags}}
	"github.com/spf13/viper"
		{{/if}}
	{{/if}}
	"github.com/spf13/cobra"
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

{{#if CLI.long}}
var {{camelT CLI.name}}Long = `{{{CLI.long}}}`
{{/if}}

{{> "flag-var.go" CLI }}

{{> "flag-init.go" CLI }}

var (
	RootCmd = &cobra.Command{

		{{#if CLI.usage}}
		Use: "{{{CLI.usage}}}",
		{{else}}
		Use: "{{{CLI.name}}}",
		{{/if}}

		{{#if CLI.short}}
		Short: "{{{CLI.short}}}",
		{{/if}}

		{{#if CLI.long}}
		Long: {{camelT CLI.name}}Long,
		{{/if}}

		{{#if CLI.persistent-prerun}}
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_persistent_prerun
			// HOFSTADTER_END   cmd_persistent_prerun
		},
		{{/if}}
		{{#if CLI.prerun}}
		PreRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PreRun {{CLI.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_prerun
			// HOFSTADTER_END   cmd_prerun
		},
		{{/if}}
		{{#unless CLI.omit-run}}
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("In {{CLI.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_run
			// HOFSTADTER_END   cmd_run
		},
		{{/unless}}
		{{#if CLI.persistent-postrun}}
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PersistentPostRun {{CLI.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_persistent_postrun
			// HOFSTADTER_END   cmd_persistent_postrun
		},
		{{/if}}
		{{#if CLI.postrun}}
		PostRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PostRun {{CLI.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_postrun
			// HOFSTADTER_END   cmd_postrun
		},
		{{/if}}
	}
)

{{/with}}

// HOFSTADTER_BELOW
