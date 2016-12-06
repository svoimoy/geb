package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.ibm.com/hofstadter-io/geb/cmd/dsl"
	"github.ibm.com/hofstadter-io/geb/cmd/info"
	"github.ibm.com/hofstadter-io/geb/cmd/project"
	"github.ibm.com/hofstadter-io/geb/engine"
	log "gopkg.in/inconshreveable/log15.v2" // logging framework
)

var (
	FlagConfigFile    string
	FlagDesignDir     string
	FlagLogLevel      string
	FlagOutputDir     string
	FlagTemplatePaths string

	logger = log.New()
)

func init() {
	//	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&FlagConfigFile, "merge-config", "", "merge a geb config file, overriding values.")
	RootCmd.PersistentFlags().StringVarP(&FlagConfigFile, "set-config", "c", "", "reset the geb config file to the file specified.")
	RootCmd.PersistentFlags().StringVarP(&FlagDesignDir, "design-dir", "d", "", "the design files directory. (default ./design)")
	RootCmd.PersistentFlags().StringVarP(&FlagLogLevel, "log-level", "l", "", "geb logging level.")
	RootCmd.PersistentFlags().StringVarP(&FlagOutputDir, "output-dir", "o", "", "the output files directory. (default ./output)")
	RootCmd.PersistentFlags().StringVarP(&FlagTemplatePaths, "template-paths", "t", "", "base templates directory. (default ./templates:~/.hofstadter/templates)")

	viper.BindPFlag("merge-config", RootCmd.PersistentFlags().Lookup("merge-config"))
	viper.BindPFlag("set-config", RootCmd.PersistentFlags().Lookup("set-config"))
	viper.BindPFlag("design-dir", RootCmd.PersistentFlags().Lookup("design-dir"))
	viper.BindPFlag("log-level", RootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("output-dir", RootCmd.PersistentFlags().Lookup("output-dir"))
	viper.BindPFlag("template-paths", RootCmd.PersistentFlags().Lookup("template-paths"))

	viper.SetDefault("design-dir", "design")
	viper.SetDefault("log-level", "warn")
	viper.SetDefault("output-dir", "output")
	viper.SetDefault("template-paths", "$HOME/.hofstadter/templates:templates")

	viper.SetConfigType("yaml")
	viper.SetConfigName("geb")

}

var (
	RootCmd = &cobra.Command{
		Use:   "geb",
		Short: "geb is a data centric code generator",
		Long: `geb is hofstadter = data + templates = profit
A data centric code generator which
combines yaml and handlebar templates
to genereate all of the things.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			read_config()
			config_logger()
		},
	}
)

func read_config() {
	viper.AddConfigPath("$HOME/.hofstadter")
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

	/*
		term_stack := log.CallerStackHandler("%+v", log.StdoutHandler)
		term_caller := log.CallerFuncHandler(log.CallerFileHandler(term_stack))
		termlog := log.LvlFilterHandler(term_level, term_caller)
	*/

	term_caller := log.CallerFuncHandler(log.CallerFileHandler(log.StdoutHandler))
	termlog := log.LvlFilterHandler(term_level, term_caller)

	//	termlog := log.LvlFilterHandler(term_level, log.StdoutHandler)
	logger.SetHandler(termlog)
	engine.SetLogger(logger)

}

func init() {
	RootCmd.AddCommand(cmd_dsl.DslCmd)
	RootCmd.AddCommand(cmd_proj.ProjectCmd)
	RootCmd.AddCommand(cmd_info.InfoCmd)
}
