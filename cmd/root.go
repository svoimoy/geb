package cmd

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ConfigPFlag        string
	DesignPFlag        string
	TemplatePathsPFlag string
	OutputPFlag        string
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
	RootCmd = &cobra.Command{
		Use:   "geb",
		Short: "geb is the Hofstadter framework CLI tool",
		Long: `Hofstadter is a Framework
for building data-centric
Platforms. geb is the tool.
`,
	}
)

// HOFSTADTER_BELOW
