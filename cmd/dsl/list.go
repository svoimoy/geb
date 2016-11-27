package cmd_dsl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listLong = `List DSLs and generators in the local library and project.

A dsl may be supplied and
one or more <generator-globs> can be appended
to limit the detail.

dsls are searched locally to remote and may by one of:
  - name   (from public library)
	- url    (of a geb-server)
	- github (of a git server)

If a project file is found,
results from the project
perspective will be listed.

note: add update flag
`

var ListCmd = &cobra.Command{
	Use:   "list [dsl] <generator-globs>",
	Short: "List DSLs and generators in the local library and project.",
	Long:  listLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list is TBD")
	},
}

func init() {
	DslCmd.AddCommand(ListCmd)
}
