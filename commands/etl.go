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
// Name:   etl
// Usage:  etl <etl-config>
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var EtlLong = `
Read file(s) or folder(s),
select and transform data,
output to file(s) or folder(s).

A more flexible, bulk processing 'geb gen adhoc'.
`

var (
	EtlInputFlag     string
	EtlInputTypeFlag string
	EtlOutputFlag    string
)

func init() {
	EtlCmd.Flags().StringVarP(&EtlInputFlag, "input", "i", "from-etl-config", "path to an input file or directory")
	viper.BindPFlag("input", EtlCmd.Flags().Lookup("input"))

	EtlCmd.Flags().StringVarP(&EtlInputTypeFlag, "input-type", "t", "yaml", "type of the data in the input file or directory")
	viper.BindPFlag("input-type", EtlCmd.Flags().Lookup("input-type"))

	EtlCmd.Flags().StringVarP(&EtlOutputFlag, "output", "o", "from-etl-config", "path to an output file or directory")
	viper.BindPFlag("output", EtlCmd.Flags().Lookup("output"))

}

var EtlCmd = &cobra.Command{

	Use: "etl <etl-config>",

	Short: "perform ETLs with geb",

	Long: EtlLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In etlCmd", "args", args)
		// Argument Parsing
		// [0]name:   etl-config
		//     help:   path to an etl config file to be used for rendering
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

		// [1]name:   sub-config
		//     help:   the name(s) of a template-config in the etl-config to use
		//     req'd:

		var subConfig []string

		if 1 < len(args) {
			subConfig = args[1:]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb etl:",
			etlConfig,
			subConfig,
		)
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(EtlCmd)
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
