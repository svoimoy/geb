package main

import (
	"fmt"
	"os"

	"{{Proj.goimport_basedir}}/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
