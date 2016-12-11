{{#with dsl.cli as |CLI| }}
package cmd

import (
  // HOFSTADTER_START import
  // HOFSTADTER_END   import
	{{#unless CLI.omit-root-run}}
	"fmt"
	{{/unless}}

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2" // logging framework

	"github.ibm.com/hofstadter-io/dotpath"
	"github.ibm.com/hofstadter-io/geb/engine"
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
		},
		{{#unless CLI.omit-root-run}}
		Run: func(cmd *cobra.Command, args []string) {
			{{> args-parse.go CLI }}

			// HOFSTADTER_START root_cmd_func
			// Do Stuff Here
			fmt.Println("dostuff")
			// HOFSTADTER_END   root_cmd_func
		},
		{{/unless}}
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

}

{{/with}}

// HOFSTADTER_BELOW
