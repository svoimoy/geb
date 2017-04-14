package commands

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	"os"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   etl
// Usage:  etl
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var EtlLong = `read file(s) or folder(s),
select and transform data,
output to file(s) or folder(s).
A more flexible and expressive
'geb gen adhoc'.
`

var EtlCmd = &cobra.Command{

	Use: "etl",

	Short: "perform ETLs with geb",

	Long: EtlLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In etlCmd", "args", args)
		// Argument Parsing
		// [0]name:   etl-config
		//     help:   a dotpath into the data to be used for rendering
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'etl-config'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var etlConfig string

		if 0 < len(args) {

			etlConfig = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb etl:",
			etlConfig,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(EtlCmd)
}

// HOFSTADTER_BELOW
