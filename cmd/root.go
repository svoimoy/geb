package cmd

import (
	"fmt"
	"strings"

	"github.com/hofstadter-io/geb/engine"
	"github.com/spf13/cobra"
)

var (
	configFile    string
	designDir     string
	templatePaths string
	outputDir     string
	generators    string
)

func init() {
	//	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "geb.yaml", "geb config file for your projectd.")
	RootCmd.PersistentFlags().StringVarP(&designDir, "design-dir", "d", "design", "the design files directory. (default ./design)")
	RootCmd.PersistentFlags().StringVarP(&templatePaths, "template-paths", "t", "~/.hofstadter/templates:./templates", "base templates directory. (default ./templates:~/.hofstadter/templates)")
	RootCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", "output", "the output files directory. (default ./output)")
	RootCmd.PersistentFlags().StringVarP(&generators, "generators", "g", "all", "which generator to run. (defaults to all found)")
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
			fmt.Println(configFile, designDir, templatePaths, outputDir, generators)
			dostuff()
		},
	}
)

func dostuff() {

	fmt.Println("geb is hofstadter = data + templates = profit")

	// Read in designs
	engine.ImportDesigns(designDir)

	// Read in templates
	t_dirs := strings.Split(templatePaths, ":")
	for _, dir := range t_dirs {
		engine.ImportTemplates(dir)
	}
}
