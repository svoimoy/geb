package system

import (
	// HOFSTADTER_START import
	"fmt"
	"path/filepath"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   DevCopyIn
// Usage:  dev-copy-in
// Parent: System
// ParentPath:

var DevCopyInLong = `Copy development files to the dot folder`

var DevCopyInCmd = &cobra.Command{
	Use:   "dev-copy-in",
	Short: "Copy development files to the dot folder",
	Long:  DevCopyInLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In DevCopyInCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		var err error
		src := "$GOPATH/src/github.ibm.com/hofstadter-io/geb"
		dest := "$HOME/.hofstadter"

		// resolve src/dest paths
		src, err = utils.ResolvePath(src)
		if err != nil {
			fmt.Println("Error:", err)
		}
		dest, err = utils.ResolvePath(dest)
		if err != nil {
			fmt.Println("Error:", err)
		}

		dirs := []string{
			"dsl",
			"library",
		}

		// do the actual copies
		for _, dir := range dirs {
			S := filepath.Join(src, dir)
			D := filepath.Join(dest, dir)
			err = utils.CopyDir(S, D)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		// HOFSTADTER_END   cmd_run
	},
}

func init() {

}

/*
Repeated Context
----------------
long: Copy development files to the dot folder
name: DevCopyIn
parent: System
path: commands.subcommands
short: Copy development files to the dot folder
usage: dev-copy-in

*/
