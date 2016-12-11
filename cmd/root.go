package cmd

import (
  // HOFSTADTER_START import
  // HOFSTADTER_END   import

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

	"github.ibm.com/hofstadter-io/dotpath"
	"github.ibm.com/hofstadter-io/geb/engine"
	"github.ibm.com/hofstadter-io/geb/cmd/system"
			"github.ibm.com/hofstadter-io/geb/cmd/view"
		"github.ibm.com/hofstadter-io/geb/cmd/gen"
			"github.ibm.com/hofstadter-io/geb/cmd/gebberish"
	
)

var (
	ConfigPFlag string
	DesignPFlag string
	TemplatePathsPFlag string
	OutputPFlag string
)


func init() {
	RootCmd.PersistentFlags().StringVarP(&ConfigPFlag, "config", "c", "geb.yaml", "A geb project config file.")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))

	RootCmd.PersistentFlags().StringVarP(&DesignPFlag, "design", "d", "design", "The design files directory.")
	viper.BindPFlag("design", RootCmd.PersistentFlags().Lookup("design"))

	RootCmd.PersistentFlags().StringVarP(&TemplatePathsPFlag, "template-paths", "t", "templates:~/.hofstadter/templates", "The search path for templates, reads from left to right, overriding along the way.")
	viper.BindPFlag("template-paths", RootCmd.PersistentFlags().Lookup("template-paths"))

	RootCmd.PersistentFlags().StringVarP(&OutputPFlag, "output", "o", "output", "The directory to output generated files to.")
	viper.BindPFlag("output", RootCmd.PersistentFlags().Lookup("output"))

}


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
		Use:   "geb",
		Short: "geb is the Hofstadter framework CLI tool",
		Long:  `Hofstadter is a Framework
for building data-centric
Platforms. geb is the tool.
`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			read_config()
			config_logger()
			logger.Debug("In PersistentPreRun Cmd", "args", args)
			// Argument Parsing
			

			// HOFSTADTER_START cmd_persistent_prerun
			// HOFSTADTER_END   cmd_persistent_prerun
		},
							}
)

func read_config() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("geb")
	viper.AddConfigPath("$HOME/.geb")
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

	system.SetLogger(logger)
			view.SetLogger(logger)
		gen.SetLogger(logger)
			gebberish.SetLogger(logger)
	

}



