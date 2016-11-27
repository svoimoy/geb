package cmd_proj

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initLong = `Initialize a new geb project.

Creates a new folder and project with the given name.
If no name is supplied, init defaults to
current working directory and uses its name.
`

var InitCmd = &cobra.Command{
	Use:   "init <name>",
	Short: "Initialize a new geb project.",
	Long:  initLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init is TBD")
	},
}

func init() {
	ProjectCmd.AddCommand(InitCmd)
}
