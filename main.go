package main

import (
	"fmt"
	"os"

	"github.ibm.com/hofstadter-io/geb/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
