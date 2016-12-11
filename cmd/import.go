package cmd

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Import
// Usage:  import <file or directory> <output file or directory>
// Parent: geb

var ImportLong = `Import other stuff into Hofstadter.

Stuff is...
  - json/jsonl, yaml, xml, protobuf, taml
  - swagger, goa.design
  - golang type definitions
`

var (
	TypeFlag string
)

func init() {
	ImportCmd.Flags().StringVarP(&TypeFlag, "type", "T", "", "The type of input data to force geb to use a certain format")
	viper.BindPFlag("type", ImportCmd.Flags().Lookup("type"))

}

var ImportCmd = &cobra.Command{
	Use: "import <file or directory> <output file or directory>",
	Aliases: []string{
		"i",
		"convert",
		"eat",
	},
	Short: "Import other stuff into Hofstadter",
	Long:  ImportLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In ImportCmd", "args", args)
		// Argument Parsing
		// [0]name:   input
		//     help:   Path to the file or folder. Can also be 'stdin'.
		//     req'd:  true
		if 0 >= len(args) {
			cmd.Usage()
			return
		}
		var input string
		if 0 < len(args) {
			input = args[0]
		}

		// [1]name:   output
		//     help:   Path to the output file or folder. Can also be 'stdout'.
		//     req'd:  true
		if 1 >= len(args) {
			cmd.Usage()
			return
		}
		var output string
		if 1 < len(args) {
			output = args[1]
		}

		// HOFSTADTER_START cmd_run
		typ := viper.GetString("type")
		fmt.Printf("In ImportCmd:  %s (%s) -> %s\n", input, typ, output)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(ImportCmd)

}

/*
Repeated Context
----------------
aliases:
- i
- convert
- eat
args:
- help: Path to the file or folder. Can also be 'stdin'.
  name: input
  required: true
  type: string
- help: Path to the output file or folder. Can also be 'stdout'.
  name: output
  required: true
  type: string
flags:
- help: The type of input data to force geb to use a certain format
  long: type
  name: Type
  short: T
  type: string
long: |
  Import other stuff into Hofstadter.

  Stuff is...
    - json/jsonl, yaml, xml, protobuf, taml
    - swagger, goa.design
    - golang type definitions
name: Import
parent: geb
path: commands
short: Import other stuff into Hofstadter
usage: import <file or directory> <output file or directory>

*/
