package cmd_info

import (
	"fmt"

	//	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var infoLong = `Print information known to the geb tool

If you are in a folder with one of the geb files,
info defaults to printing information about that component.
Otherwise it prints information known to the geb tool
from its default location (~/.hofstadter).

See the sub-commands for additional options.
`

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print info known to the geb tool.",
	Long:  infoLong,
	Run: func(cmd *cobra.Command, args []string) {

		do_info_cmd(args)

	},
}

// This function is reached when an
// info sub-command is not run
func do_info_cmd(args []string) {

	var (
		err error
	)

	// What file are we dealing with?
	// That is how we decide the default behavior
	file := look_for_file()
	switch file {
	case "geb.yml", "geb.yaml":
		err = project_info(file, args)

	case "geb-dsl.yml", "geb-dsl.yaml":
		err = dsl_info(file, args)

	case "geb-gen.yml", "geb-gen.yaml":
		err = gen_info(file, args)

	default:
		err = geb_info(args)
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	return

}
