package cmd

// this command is for importing
// various file types to geb designs

import (
	"fmt"

	"github.com/spf13/cobra"
)

var importLong = `geb import [type] [file]
imports a file type to a geb design file
to a folder named 'imports' in the output dir.

Intended file formats are:
 - goa.design
 - types
	 - protobuf
	 - json
	 - xml
	 - golang
 - apis
	 - swagger
 - dsls
   - kubernetes
`

var ImportCmd = &cobra.Command{
	Use:   "import [type] [file]",
	Short: "Import another file type to geb design.",
	Long:  generateLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import is TBD")
	},
}

func init() {
	RootCmd.AddCommand(ImportCmd)
}
