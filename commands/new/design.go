package new

import (
	// HOFSTADTER_START import
	"fmt"

	libnew "github.com/hofstadter-io/geb/lib/new"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   design
// Usage:  design <dsl-name> [name]
// Parent: new

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var DesignLong = `Initialize a new geb design for a dsl. If you do not provide a name, the current directory name will be used.`

var DesignCmd = &cobra.Command{

	Use: "design <dsl-name> [name]",

	Short: "Initialize a new geb design for a dsl.",

	Long: DesignLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In designCmd", "args", args)
		// Argument Parsing
		// [0]name:   dsl
		//     help:   The relative path for the dsl/generator, relative from the dsl paths specified in geb.yaml
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'dsl'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var dsl string

		if 0 < len(args) {

			dsl = args[0]
		}

		// [1]name:   gen
		//     help:   The relative path for the dsl/generator, relative from the dsl paths specified in geb.yaml
		//     req'd:  true
		if 1 >= len(args) {
			fmt.Println("missing required argument: 'gen'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var gen string

		if 1 < len(args) {

			gen = args[1]
		}

		// [2]name:   name
		//     help:   The name for the dsl
		//     req'd:

		var name string

		if 2 < len(args) {

			name = args[2]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb new design:",
			dsl,
			name,
		)
		err := libnew.NewDesign(dsl, gen, name)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
