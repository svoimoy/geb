package cmd_info

import (
	"fmt"
	"github.com/pkg/errors"

	// "github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var gebLong = `Print information known to the geb tool.

This subcommand prints information known to the geb tool
from its default location (~/.hofstadter).

See the sub-commands for additional options.
`

var GebCmd = &cobra.Command{
	Use:   "geb <args>... <field>...",
	Short: "Print information known to the geb tool.",
	Long:  gebLong,
	Run: func(cmd *cobra.Command, args []string) {

		err := geb_info(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		return

	},
}

func init() {
	InfoCmd.AddCommand(GebCmd)
}

func geb_info(args []string) error {

	return errors.New("Not implemented yet")
}
