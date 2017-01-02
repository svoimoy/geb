package cmd

import (
	// HOFSTADTER_START import
	"fmt"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/cmd/gen"
)

// Tool:   geb
// Name:   Gen
// Usage:  gen
// Parent: geb

var GenLong = `Generate a project from its working directory.`

var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a project.",
	Long:  GenLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In GenCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		filename := "geb.yaml"

		err := engine.GenerateProject(filename, args)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(GenCmd)

	gen.SetLogger(logger)
	GenCmd.AddCommand(gen.FileCmd)

}

// HOFSTADTER_BELOW
