package system

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Update
// Usage:  update
// Parent: System
// ParentPath:

var UpdateLong = `Update the geb library DSLs, designs, and other files in the dot folder.`

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the geb library and dot folder",
	Long:  UpdateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In UpdateCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		// HOFSTADTER_END   cmd_run
	},
}

func init() {

}

// HOFSTADTER_BELOW
