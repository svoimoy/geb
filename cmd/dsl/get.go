package cmd_dsl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getLong = `Download DSLs and generators to the local machine.

dsls are searched locally to remote and may by one of:
  - name   (from public library)
	- url    (of a geb-server)
	- github (of a git server)

note: add update flag
`

var GetCmd = &cobra.Command{
	Use:   "get [dsl] <generator-globs>",
	Short: "Download DSLs and generators to the local machine.",
	Long:  getLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get is TBD")
	},
}

func init() {
	DslCmd.AddCommand(GetCmd)
}
