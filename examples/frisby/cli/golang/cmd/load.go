package cmd

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import


	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   frisby
// Name:   Load
// Usage:  load <file> <host>
// Parent: frisby

var LoadLong = `frisby is an API testing and thrashing toolset`


var (
	WorkersFlag int
)


func init() {
	LoadCmd.Flags().StringVarP(&WorkersFlag, "workers", "w", "8", "the number of workers")
	viper.BindPFlag("workers", LoadCmd.Flags().Lookup("workers"))
	
}

var LoadCmd = &cobra.Command {
	Use: "load <file> <host>",
	Short: "Run a frisby test file, loading a host.",
	Long: LoadLong,
	Run: func(cmd *cobra.Command, args []string) {
		// HOFSTADTER_START cmd_run
		fmt.Println("In LoadCmd")
		// HOFSTADTER_END   cmd_run
	},
}


func init() {
	RootCmd.AddCommand(LoadCmd)

}

