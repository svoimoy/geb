{{#with dsl.cli as |CLI| }}
package cmd

import (
  // HOFSTADTER_START import
  // HOFSTADTER_END   import

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

{{> "flag-var.go" CLI }}

{{> "flag-init.go" CLI }}

var (
	RootCmd = &cobra.Command{
		Use:   "{{ CLI.name }}",
		Short: "{{ CLI.short }}",
		Long:  `{{ CLI.long }}`,
		{{#if CLI.persistent-prerun}}
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_persistent_prerun
			// HOFSTADTER_END   cmd_persistent_prerun
		},
		{{/if}}
		{{#if CLI.prerun}}
		PreRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PreRun {{RC.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_prerun
			// HOFSTADTER_END   cmd_prerun
		},
		{{/if}}
		{{#unless CLI.omit-run}}
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("In {{RC.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_run
			// HOFSTADTER_END   cmd_run
		},
		{{/unless}}
		{{#if CLI.persistent-postrun}}
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PersistentPostRun {{RC.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_persistent_postrun
			// HOFSTADTER_END   cmd_persistent_postrun
		},
		{{/if}}
		{{#if CLI.postrun}}
		PostRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PostRun {{RC.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_postrun
			// HOFSTADTER_END   cmd_postrun
		},
		{{/if}}
	}
)

{{/with}}

// HOFSTADTER_BELOW
