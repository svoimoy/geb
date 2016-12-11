package main

import (
	"fmt"
	"os"

	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/' )}}}/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
