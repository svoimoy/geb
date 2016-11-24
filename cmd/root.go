package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile    string
	designDir     string
	templatePaths string
	outputDir     string
	generators    string
	verbose       bool
)

func init() {
	viper.SetConfigType("yaml")

	//	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "geb config file for your projectd.")
	RootCmd.PersistentFlags().StringVarP(&designDir, "design-dir", "d", "", "the design files directory. (default ./design)")
	RootCmd.PersistentFlags().StringVarP(&templatePaths, "template-paths", "t", "", "base templates directory. (default ./templates:~/.hofstadter/templates)")
	RootCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", "", "the output files directory. (default ./output)")
	RootCmd.PersistentFlags().StringVarP(&generators, "generators", "g", "", "which generator to run. (defaults to all found)")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print out verbose messages to see whats giong on.")

	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("design-dir", RootCmd.PersistentFlags().Lookup("design-dir"))
	viper.BindPFlag("template-paths", RootCmd.PersistentFlags().Lookup("template-paths"))
	viper.BindPFlag("output-dir", RootCmd.PersistentFlags().Lookup("output-dir"))
	viper.BindPFlag("generators", RootCmd.PersistentFlags().Lookup("generators"))

	viper.SetDefault("config", "geb.yaml")
	viper.SetDefault("design-dir", "design")
	viper.SetDefault("template-paths", "~/.hofstadter/templates:./templates")
	viper.SetDefault("output-dir", "output")
	viper.SetDefault("generators", "all")

	viper.SetConfigName(viper.Get("config").(string))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if err.Error() == "open : no such file or directory" { // Handle errors reading the config file
			if verbose {
				fmt.Println("No 'geb.yaml' file found. Use 'geb project init' to create one.")
			}
		} else {
			fmt.Println(err)
		}
	}
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
