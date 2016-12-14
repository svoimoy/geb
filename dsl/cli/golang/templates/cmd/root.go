{{#with dsl.cli as |CLI| }}
package cmd

import (
  // HOFSTADTER_START import
  // HOFSTADTER_END   import

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

	"github.ibm.com/hofstadter-io/dotpath"
	"github.ibm.com/hofstadter-io/geb/engine"
	{{#each CLI.commands as |Cmd|}}
	{{#if Cmd.subcommands}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/{{lower Cmd.name}}"
	{{/if}}
	{{/each}}
)

{{> "flag-var.go" CLI }}

{{> "flag-init.go" CLI }}

var (
	FlagMergeConfigFile    string
	FlagSetConfigFile    string
	FlagLogLevel    string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&FlagLogLevel, "log-level", "l", "", "reset the geb config file to the file specified.")
	RootCmd.PersistentFlags().StringVar(&FlagMergeConfigFile, "merge-config", "", "merge a geb config file, overriding values.")
	RootCmd.PersistentFlags().StringVar(&FlagSetConfigFile, "set-config", "", "reset the geb config file to the file specified.")

	viper.BindPFlag("log-level", RootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("merge-config", RootCmd.PersistentFlags().Lookup("merge-config"))
	viper.BindPFlag("set-config", RootCmd.PersistentFlags().Lookup("set-config"))

	viper.SetDefault("log-level", "warn")

}

var (
	logger = log.New()
	
	RootCmd = &cobra.Command{
		Use:   "{{ CLI.name }}",
		Short: "{{ CLI.short }}",
		Long:  `{{ CLI.long }}`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			read_config()
			config_logger()
			logger.Debug("In PersistentPreRun {{RC.name}}Cmd", "args", args)
			{{> args-parse.go CLI }}

			// HOFSTADTER_START cmd_persistent_prerun
			// HOFSTADTER_END   cmd_persistent_prerun
		},
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

func read_config() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("{{CLI.name}}")
	viper.AddConfigPath("$HOME/.{{CLI.name}}")
	viper.MergeInConfig()
	viper.AddConfigPath(".")
	viper.MergeInConfig()

	cfg := viper.GetString("add-config")
	if cfg != "" {
		viper.SetConfigFile(cfg)
		viper.MergeInConfig()
	}

	cfg = viper.GetString("set-config")
	if cfg != "" {
		viper.SetConfigFile(cfg)
		viper.ReadInConfig()
	}
}

func config_logger() {
	level := viper.GetString("log-level")
	term_level, err := log.LvlFromString(level)
	if err != nil {
		panic(err)
	}

	// term_stack := log.CallerStackHandler("%+v", log.StdoutHandler)
	// term_caller := log.CallerFuncHandler(log.CallerFileHandler(term_stack))
	// termlog := log.LvlFilterHandler(term_level, term_caller)

	termlog := log.LvlFilterHandler(term_level, log.StdoutHandler)
	logger.SetHandler(termlog)
	dotpath.SetLogger(logger)
	engine.SetLogger(logger)

	{{#each CLI.commands as |Cmd|}}
	{{#if Cmd.subcommands}}
	{{lower Cmd.name}}.SetLogger(logger)
	{{/if}}
	{{/each}}

	// HOFSTADTER_START config_logger
	// HOFSTADTER_END   config_logger
}


{{/with}}

