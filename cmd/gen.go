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





var GenCmd = &cobra.Command {
	Use: "gen",
	Short: "Generate a project.",
	Long: GenLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		gen.SetLogger(logger)
		logger.Debug("In PersistentPreRun GenCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
	
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

	GenCmd.AddCommand(gen.FileCmd)

	gen.SetLogger(logger)
}


/*
Repeated Context
----------------
long: Generate a project from its working directory.
name: Gen
parent: geb
path: commands
short: Generate a project.
subcommands:
- args:
  - help: Path to the input design file.
    name: designFile
    required: true
    type: string
  - help: Path to the template file.
    name: templateFile
    required: true
    type: string
  - help: Path to the output file. Can also be 'stdout'.
    name: outputFile
    required: true
    type: string
  long: Generate a file from design and a template.
  name: File
  parent: Gen
  path: commands.subcommands
  short: Generate a file.
  usage: file <designFile> <templateFile> <outputFile>
usage: gen

*/
