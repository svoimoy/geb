package cmd

import (
	"fmt"
	// "io/ioutil"

	"github.com/hofstadter-io/geb/engine/project"
	"github.com/spf13/cobra"
)

var DevCmd = &cobra.Command{
	Use:   "dev [args]",
	Short: "run development code",
	Long:  "run the development code command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running Dev Command\n\n")

		filename := "geb.yaml"
		if len(args) == 1 {
			filename = args[0]
		}

		cfg, err := project.ReadConfigFile(filename)
		if err != nil {
			fmt.Println("Error reading project config "+filename+":", err)
			return
		}
		fmt.Printf("config:\n%+v\n\n:", cfg)
	},
}

func init() {
	RootCmd.AddCommand(DevCmd)
}
