package cmd_info

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var projectLong = `Print information about a project.

If you are in a folder with one of the geb files,
info defaults to printing information about that component.
Otherwise it prints information known to the geb tool
from its default location (~/.hofstadter).

See the sub-commands for additional options.
`

var ProjectCmd = &cobra.Command{
	Use:   "proj <field>...",
	Short: "Print information about a project.",
	Long:  projectLong,
	Run: func(cmd *cobra.Command, args []string) {

		var (
			data interface{}
			err  error
		)

		file := look_for_file()
		switch file {
		case "geb.yml", "geb.yaml":
			data, err = project_info(file, args)

		default:
			fmt.Println("No project file found in the current directory.")
			return
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
	InfoCmd.AddCommand(ProjectCmd)
}

func project_info(filename string, args []string) (interface{}, error) {

	return nil, errors.New("Not implemented yet")
}