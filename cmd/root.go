package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	FlagConfigFile    string
	FlagDesignDir     string
	FlagTemplatePaths string
	FlagOutputDir     string
	LOUD              bool
)

func init() {
	//	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&FlagConfigFile, "config", "c", "", "geb config file for your projectd.")
	RootCmd.PersistentFlags().StringVarP(&FlagDesignDir, "design-dir", "d", "", "the design files directory. (default ./design)")
	RootCmd.PersistentFlags().StringVarP(&FlagTemplatePaths, "template-paths", "t", "", "base templates directory. (default ./templates:~/.hofstadter/templates)")
	RootCmd.PersistentFlags().StringVarP(&FlagOutputDir, "output-dir", "o", "", "the output files directory. (default ./output)")
	RootCmd.PersistentFlags().BoolVarP(&LOUD, "verbose", "v", false, "Print out verbose messages to see whats giong on.")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("design-dir", RootCmd.PersistentFlags().Lookup("design-dir"))
	viper.BindPFlag("template-paths", RootCmd.PersistentFlags().Lookup("template-paths"))
	viper.BindPFlag("output-dir", RootCmd.PersistentFlags().Lookup("output-dir"))

	viper.SetConfigType("yaml")
	viper.SetConfigName("geb")
	viper.AddConfigPath(".")

	viper.SetDefault("design-dir", "design")
	viper.SetDefault("template-paths", "$HOME/.hofstadter/templates:templates")
	viper.SetDefault("output-dir", "output")
	viper.SetDefault("verbose", false)

}

var (
	RootCmd = &cobra.Command{
		Use:   "geb",
		Short: "geb is a data centric code generator",
		Long: `A data centric code generator which
combines yaml and handlebar templates
to genereate all of the codes.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			// dostuff()
			fmt.Println("geb is hofstadter = data + templates = profit")
			cmd.Usage()
		},
	}
)
