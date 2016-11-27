package cmd_proj

import (
	// "fmt"

	"github.com/spf13/cobra"
)

var projectLong = `project
Manage geb projects.
`

var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage geb projects.",
	Long:  projectLong,
}
