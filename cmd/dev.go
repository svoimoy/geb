package cmd

import (
	"fmt"

	"github.com/hofstadter-io/geb/engine"
	"github.com/spf13/cobra"
)

var DevCmd = &cobra.Command{
	Use:   "dev [args]",
	Short: "run development code",
	Long:  "run the development code command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running Dev Command\n----------------")
		engine.Dev(args)
	},
}

func init() {
	RootCmd.AddCommand(DevCmd)
}
