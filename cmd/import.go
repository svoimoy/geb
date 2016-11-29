package cmd

// this command is for importing
// various file types to geb designs

import (
	"fmt"

	"github.com/spf13/cobra"
)

var importLong = `imports a file(s)/folder(s) type to a geb design file(s)/folder(s)
to a folder named 'imports' in the output dir.

Intended mportable formats are:

 - swagger
 - json
 - protobuf
 - golang
 - xml
 - sql table create
 - goa.design
`

var ImportCmd = &cobra.Command{
	Use:   "import [type] [file(s)/folder(s)]",
	Short: "Import another file type into geb design.",
	Long:  generateLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import is TBD")
	},
}

func init() {
	RootCmd.AddCommand(ImportCmd)
}
