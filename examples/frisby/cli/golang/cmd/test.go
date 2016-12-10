package cmd

import (
	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"
)

// Tool:   frisby
// Name:   Test
// Usage:  test <file> <host>
// Parent: frisby

var TestLong = `frisby is an API testing and thrashing toolset`





var TestCmd = &cobra.Command {
	Use: "test <file> <host>",
	Short: "Run a frisby test file against a host.",
	Long: TestLong,
	Run: func(cmd *cobra.Command, args []string) {
		// HOFSTADTER_START cmd_run
		fmt.Println("In TestCmd")
		// HOFSTADTER_END   cmd_run
	},
}


func init() {
	RootCmd.AddCommand(TestCmd)

}

