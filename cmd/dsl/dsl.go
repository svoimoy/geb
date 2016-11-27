package cmd_dsl

// This command is intended for managing geb packages
// mainly related to the community contributed content
// should support:
//   - hofstadter.io geb-hub library and namespace/hosted
//   - arbitrary geb-hub server
//   - github and friends
//
// i.e. Official content lacks a domain name prefix
// geb-hosted, self-hosted, and github and friends
// have the prefix
// autogen geb-hosted from a github repo

import (
	// "fmt"

	"github.com/spf13/cobra"
)

var dslLong = `Manage geb DSLs and generators.
`

var DslCmd = &cobra.Command{
	Use:   "dsl",
	Short: "Manage geb DSLs and generators.",
	Long:  dslLong,
}
