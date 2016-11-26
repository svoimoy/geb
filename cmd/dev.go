package cmd

import (
	"fmt"
	// "io/ioutil"

	"github.com/hofstadter-io/geb/engine"
	"github.com/spf13/cobra"
)

var DevCmd = &cobra.Command{
	Use:   "dev [args]",
	Short: "run development code",
	Long:  "run the development code command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running Dev Command\n")
		engine.SayHello()
	},
}

func init() {
	RootCmd.AddCommand(DevCmd)
}
