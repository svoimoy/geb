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
// Name:   generator
// Usage:  generator <name>
// Parent: new

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var GeneratorLong = `Initialize a new geb generator. If you do not provide a name, the current directory name will be used. The dsl type will be infered by looking recursively up the parent directories until a geb-dsl found.`

var GeneratorCmd = &cobra.Command{

	Use: "generator <name>",

	Aliases: []string{
		"gen",
	},

	Short: "Initialize a new geb generator.",

	Long: GeneratorLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In generatorCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:   The name for the generator
		//     req'd:

		var name string

		if 0 < len(args) {

			name = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb new generator:",
			name,
		)
		libnew.NewGenerator(name)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
