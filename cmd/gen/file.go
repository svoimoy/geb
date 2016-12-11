package gen

import (
	// HOFSTADTER_START import
	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"fmt"

	
	"github.com/spf13/cobra"

)

// Tool:   geb
// Name:   File
// Usage:  file <designFile> <templateFile> <outputFile>
// Parent: Gen
// ParentPath: 

var FileLong = `Generate a file from design and a template.`





var FileCmd = &cobra.Command {
	Use: "file <designFile> <templateFile> <outputFile>",
	Short: "Generate a file.",
	Long: FileLong,
		
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In FileCmd", "args", args)
		// Argument Parsing
		// [0]name:   designFile
		//     help:   Path to the input design file.
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("not enough args supplied")
			return
		}
		var designFile string
		if 0 < len(args) {
			designFile = args[0]
		}
		
		// [1]name:   templateFile
		//     help:   Path to the template file.
		//     req'd:  true
		if 1 >= len(args) {
			fmt.Println("not enough args supplied")
			return
		}
		var templateFile string
		if 1 < len(args) {
			templateFile = args[1]
		}
		
		// [2]name:   outputFile
		//     help:   Path to the output file. Can also be 'stdout'.
		//     req'd:  true
		if 2 >= len(args) {
			fmt.Println("not enough args supplied")
			return
		}
		var outputFile string
		if 2 < len(args) {
			outputFile = args[2]
		}
		
		

		// HOFSTADTER_START cmd_run
		err := engine.GenerateFile(designFile, templateFile, outputFile)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		// HOFSTADTER_END   cmd_run
	},
		}


func init() {

}


/*
Repeated Context
----------------
args:
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

*/
