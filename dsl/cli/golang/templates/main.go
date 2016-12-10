package main

import (
	"fmt"
	"os"

	"{{{trimprefix gen_basedir (concat2 ENV.GOPATH '/src/' )}}}/cmd"
	"{{{trimprefix gen_basedir (concat2 (getenv 'GOPATH') '/src/' )}}}/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
