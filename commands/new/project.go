package new

import (
	// HOFSTADTER_START import
	"fmt"

	libnew "github.com/hofstadter-io/geb/lib/new"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   project
// Usage:  project [name]
// Parent: new

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var ProjectLong = `Initialize a new geb project. If you do not provide a name, the current directory name will be used.`

var ProjectCmd = &cobra.Command{

	Use: "project [name]",

	Aliases: []string{
		"proj",
	},

	Short: "Initialize a new geb project.",

	Long: ProjectLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In projectCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:   The name for the project
		//     req'd:

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb new project:",
			name,
		)
		libnew.NewProject(name)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
