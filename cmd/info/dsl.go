package cmd_info

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var dslLong = `Print information about DSLs.

If you are in a folder with one of the geb files,
info defaults to printing information about that component.
Otherwise it prints information known to the geb tool
from its default location (~/.hofstadter).

See the sub-commands for additional options.
`

var DslCmd = &cobra.Command{
	Use:   "dsl <name> <field>...",
	Short: "Print information about DSLs.",
	Long:  dslLong,
	Run: func(cmd *cobra.Command, args []string) {

		var (
			data interface{}
			err  error
		)

		file := look_for_file()
		switch file {
		case "geb.yml", "geb.yaml":
			args = append([]string{"dsl"}, args...)
			data, err = project_info(file, args)

		case "geb-dsl.yml", "geb-dsl.yaml":
			data, err = dsl_info(file, args)

		default:
			data, err = dsl_info("", args)
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%# v", pretty.Formatter(data))
		return
	},
}

func init() {
	InfoCmd.AddCommand(DslCmd)
}

func dsl_info(filename string, args []string) (interface{}, error) {

	return nil, errors.New("Not implemented yet")
}
