package cmd_proj

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cloneLong = `Clone an existing geb project.

Clones the project into the working directory.

Projects are searched locally to remote and may by one of:
  - path   (on the local machine)
  - name   (from public library)
	- url    (of a geb-server)
	- github (of a git server)
`

var CloneCmd = &cobra.Command{
	Use:   "clone <name-or-url>",
	Short: "Clone an existing geb project.",
	Long:  cloneLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clone is TBD")
	},
}

func init() {
	ProjectCmd.AddCommand(CloneCmd)
}
