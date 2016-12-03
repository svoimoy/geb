package cmd

import (
	"fmt"

	"github.com/hofstadter-io/geb/engine"
	"github.com/spf13/cobra"
)

var fileLong = `Renders a single output from
a design file and a template file.
If no output file is specified,
the rendered template will go to stdout.`

var FileCmd = &cobra.Command{
	Use:   "file <design-file> <template-file> [output-file]",
	Short: "render a single file from one design and template file each.",
	Long:  fileLong,
	Run: func(cmd *cobra.Command, args []string) {
		// Check number of args, should be 2 or 3
		if len(args) < 2 || 3 < len(args) {
			cmd.Usage()
		}

		design := args[0]
		template := args[1]
		output := "stdout"
		if len(args) == 3 {
			output = args[2]
		}

		err := engine.GenerateFile(design, template, output)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	GenerateCmd.AddCommand(FileCmd)
}
