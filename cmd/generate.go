package cmd

import (
	"strings"

	"github.com/hofstadter-io/geb/engine"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateLong = `geb generate processes the current directory.
It reads the design and template directories,
runs the specified generators,
and outputs the rendered files.`

var GenerateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Run the geb generator in the current directory.",
	Long:  generateLong,
	Run: func(cmd *cobra.Command, args []string) {
		// Check number of args, should be 2 or 3
		if len(args) != 0 {
			cmd.Help()
		}

		// Read in designs
		engine.ImportDesignFolder(viper.Get("design-dir").(string))

		// Read in templates
		t_dirs := strings.Split(viper.Get("template-paths").(string), ":")
		for _, dir := range t_dirs {
			engine.ImportTemplateFolder(dir)
		}

		outdir := viper.Get("output-dir").(string)
		geners := viper.Get("generators").(string)
		generators := strings.Split(geners, ",")
		engine.Generate(generators, outdir)

	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
}
