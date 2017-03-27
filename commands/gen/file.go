package gen
// package subcommands

import (
	// HOFSTADTER_START import
	"fmt"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	
	"github.com/spf13/cobra"

)

// Tool:   geb
// Name:   file
// Usage:  file <designFile> <templateFile> <outputFile>
// Parent: gen

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init


var FileLong = `Generate a file from design and a template.`






var FileCmd = &cobra.Command {
	Use: "file <designFile> <templateFile> <outputFile>",
	Short: "Generate a file.",
	Long: FileLong,
		
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In fileCmd", "args", args)
		// Argument Parsing
		// [0]name:   designFile
		//     help:   Path to the input design file.
		//     req'd:  true
		if 0 >= len(args) {
			cmd.Usage()
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
			cmd.Usage()
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
			cmd.Usage()
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



// HOFSTADTER_BELOW
