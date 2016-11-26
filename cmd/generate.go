package cmd

import (
	"fmt"
	// "strings"

	// "github.com/hofstadter-io/geb/engine"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateLong = `geb generate processes the current directory.
It reads the design and template directories,
runs the specified generators,
and outputs the rendered files.
You can specify which generators should be
run by default in your geb.yaml file.
Specify one or more generators
as arguments to override those defaults.`

var GenerateCmd = &cobra.Command{
	Use:   "gen [generator]...",
	Short: "Run the geb generator in the current directory.",
	Long:  generateLong,
	PreRun: func(cmd *cobra.Command, args []string) {
		cfg := viper.Get("config").(string)
		if cfg != "" {
			fmt.Println("using config file: ", cfg)
			viper.SetConfigFile(cfg)
		}

		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {
			if err.Error() == "open : no such file or directory" { // Handle errors reading the config file
				if LOUD {
					fmt.Println("No config file found. Use 'geb project init' to create one.")
				}
			} else {
				fmt.Println(err)
			}
		} else {
			if LOUD {
				fmt.Println(cfg, "file found.")
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		/*
			var geners []string

			// Check number of args, should be no args
			if len(args) == 0 {
				geners = viper.GetStringSlice("generators")
			} else {
				geners = args
			}

				// Read in designs
				d_dir := viper.GetString("design-dir")
				fmt.Println("Loading designs from:", d_dir)
				err := engine.ImportDesignFolder(d_dir)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				// Read in templates
				fmt.Println("Loading Templates:")
				t_dirs := strings.Split(viper.Get("template-paths").(string), ":")
				for _, dir := range t_dirs {
					fmt.Println("     ", dir)
					err := engine.ImportTemplateFolder(dir)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
				}

				// make rendering plans
				plans, err := engine.MakeRenderingPlans(geners, engine.DESIGN, engine.TEMPLATES)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				// Render the files
				outdir := viper.Get("output-dir").(string)
				fmt.Println("Rendering files to:", outdir)
				err = engine.RenderPlans(plans, outdir)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
		*/

	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
}
