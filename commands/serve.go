package commands

import (
	// HOFSTADTER_START import
	"fmt"

	"github.com/hofstadter-io/geb/lib/serve"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	"os"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   serve
// Usage:  serve <server-config>
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var ServeLong = `"similar to the 'geb adhoc' and 'geb etl' commands
now served up as a RESTful API.

See the server docs at: ...tbd..."
`

var ServeCmd = &cobra.Command{

	Use: "serve <server-config>",

	Short: "serve templates and ETL pipelines with geb",

	Long: ServeLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In serveCmd", "args", args)
		// Argument Parsing
		// [0]name:   server-config
		//     help:   a dotpath into the server configuration file
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'server-config'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var serverConfig string

		if 0 < len(args) {

			serverConfig = args[0]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb serve:",
			serverConfig,
		)

		err := serve.Run()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(ServeCmd)
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
