package system

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   update
// Usage:  update
// Parent: system

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var UpdateLong = `Update the geb library DSLs, designs, and other files in the dot folder.`

var UpdateCmd = &cobra.Command{

	Use: "update",

	Short: "Update the geb library and dot folder",

	Long: UpdateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In updateCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		fmt.Println("geb system update will be updated when the website goes live.")
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
