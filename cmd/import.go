package cmd

// this command is for converting
// various file types to geb designs

import (
	"fmt"

	"github.com/spf13/cobra"
)

var convertLong = `converts a file(s)/folder(s) type to a geb design file(s)/folder(s)
to a folder named 'converts' in the output dir.

Intended mportable formats are:

 - swagger
 - json
 - protobuf
 - golang
 - xml
 - sql table create
 - goa.design
`

var ConvertCmd = &cobra.Command{
	Use:   "convert [type] [file(s)/folder(s)]",
	Short: "Convert another file type into geb design.",
	Long:  generateLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert is TBD")
	},
}

func init() {
	RootCmd.AddCommand(ConvertCmd)
}
