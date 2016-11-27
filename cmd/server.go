package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverLong = `Run the geb server.

Environment settings can be declared in
the server config section and used here.

Watch mode enables the server to monitor files,
making changes and reloading as needed.
`

var FlagWatch bool

var ServerCmd = &cobra.Command{
	Use:   "server <env>",
	Short: "Run the geb server.",
	Long:  serverLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server is TBD")
	},
}

func init() {
	ServerCmd.Flags().BoolVarP(&FlagWatch, "watch", "w", false, "Run the server in watching mode")

	RootCmd.AddCommand(ServerCmd)
}
