package commands

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import

	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   etl-serve
// Usage:  etl-serve <etl-config>
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var EtlServeLong = `similar to the 'geb etl' command
and served up in as a RESTful API.
`

var (
	hostFlag string
	portFlag string
)

func init() {
	EtlServeCmd.Flags().StringVarP(&hostFlag, "host", "h", "localhost", "host to run the server as")
	viper.BindPFlag("host", EtlServeCmd.Flags().Lookup("host"))

	EtlServeCmd.Flags().StringVarP(&portFlag, "port", "p", "1110", "port to run the server on")
	viper.BindPFlag("port", EtlServeCmd.Flags().Lookup("port"))

}

var EtlServeCmd = &cobra.Command{

	Use: "etl-serve <etl-config>",

	Short: "serve ETLs with geb",

	Long: EtlServeLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In etl-serveCmd", "args", args)
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
		fmt.Println("geb etl-serve:",
			etlConfig,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(EtlServeCmd)
}

// HOFSTADTER_BELOW
