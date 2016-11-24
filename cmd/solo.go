package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/hofstadter-io/geb/engine"
	"github.com/spf13/cobra"
)

var soloLong = `Solo renders a single output from
a design file and a template file.
If no output file is specified,
the rendered template will go to stdout.`

var SoloCmd = &cobra.Command{
	Use:   "solo <design> <template> [output-file]",
	Short: "render a single file from one design and template file each.",
	Long:  soloLong,
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

		// fmt.Printf("%s + %s => %s\n", design, template, output)

		err := engine.ImportDesignFile(design)
		if err != nil {
			fmt.Println("Error reading design:", err)
			cmd.Help()
		}

		err = engine.ImportTemplateFile(template)
		if err != nil {
			fmt.Println("Error reading template:", err)
			cmd.Help()
		}

		result, err := engine.RenderTemplate(template, engine.DESIGN)
		if err != nil {
			fmt.Println("Error while rendering:", err)
			cmd.Help()
		}

		if output == "stdout" {
			fmt.Println(result)
			return
		}

		err = ioutil.WriteFile(output, []byte(result), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			cmd.Help()
		}

	},
}

func init() {
	RootCmd.AddCommand(SoloCmd)
}
