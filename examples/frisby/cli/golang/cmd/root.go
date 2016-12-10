package cmd

import (
  // HOFSTADTER_START import
  // HOFSTADTER_END   import

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2" // logging framework
)




var (
	FlagMergeConfigFile    string
	FlagSetConfigFile    string
)

func init() {
	RootCmd.PersistentFlags().StringVar(&FlagMergeConfigFile, "merge-config", "", "merge a geb config file, overriding values.")
	RootCmd.PersistentFlags().StringVar(&FlagSetConfigFile, "set-config", "", "reset the geb config file to the file specified.")

	viper.BindPFlag("merge-config", RootCmd.PersistentFlags().Lookup("merge-config"))
	viper.BindPFlag("set-config", RootCmd.PersistentFlags().Lookup("set-config"))
}

var (
	logger = log.New()
	
	RootCmd = &cobra.Command{
		Use:   "frisby",
		Short: "frisby is an API testing and thrashing toolset",
		Long:  `frisby is an API testing and thrashing toolset`,
	}
)

func read_config() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("frisby")
	viper.AddConfigPath("$HOME/.frisby")
	viper.ReadInConfig()
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
		viper.MergeInConfig()
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

}


// HOFSTADTER_BELOW

// HOFSTADTER_BELOW
