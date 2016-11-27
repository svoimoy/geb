package cmd_dsl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateLong = `update DSLs and generators in the local project or machine.

dsls are searched locally to remote and may by one of:
  - name   (from public library)
	- url    (of a geb-server)
	- github (of a git server)

Note: add global flag
`

var UpdateCmd = &cobra.Command{
	Use:   "update [dsl] <generator-globs>",
	Short: "Update DSLs and generators to the local project or machine.",
	Long:  updateLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update is TBD")
	},
}

func init() {
	DslCmd.AddCommand(UpdateCmd)
}
