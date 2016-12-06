package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.ibm.com/hofstadter-io/geb/engine"
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
	Run: func(cmd *cobra.Command, args []string) {

		filename := "geb.yaml"

		err := engine.GenerateProject(filename, args)
		if err != nil {
			fmt.Println("Error:", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
}
