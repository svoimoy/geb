package cmd_dsl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addLong = `Add DSLs and generators to the current project.

dsls are searched locally to remote and may by one of:
  - name   (from public library)
	- url    (of a geb-server)
	- github (of a git server)
`

var AddCmd = &cobra.Command{
	Use:   "add [dsl] <generator-globs>",
	Short: "Add DSLs and generators to the current project.",
	Long:  addLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add is TBD")
	},
}

func init() {
	DslCmd.AddCommand(AddCmd)
}
