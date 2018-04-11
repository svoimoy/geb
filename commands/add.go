package commands

import (
	// HOFSTADTER_START import
	"fmt"
	"path/filepath"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	"os"

	"github.com/hofstadter-io/geb/lib/add"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   add
// Usage:  add [sub-command] [git-repo-url]
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var AddLong = `Add installs a design, dsl, or generator to the current project.

See [docs link...] for more information.
`

var (
	AddGlobalPFlag bool

	AddBranchPFlag string

	AddTagPFlag string

	AddCommitPFlag string

	AddSubmodulePFlag bool
)

func init() {
	AddCmd.PersistentFlags().BoolVarP(&AddGlobalPFlag, "global", "g", false, "add the package to the global context in ~/.geb/dsl/... Ignored by designs")
	viper.BindPFlag("global", AddCmd.PersistentFlags().Lookup("global"))

	AddCmd.PersistentFlags().StringVarP(&AddBranchPFlag, "branch", "b", "master", "The branch to check out.")
	viper.BindPFlag("branch", AddCmd.PersistentFlags().Lookup("branch"))

	AddCmd.PersistentFlags().StringVarP(&AddTagPFlag, "tag", "t", "", "The tag to check out. Overrides branch.")
	viper.BindPFlag("tag", AddCmd.PersistentFlags().Lookup("tag"))

	AddCmd.PersistentFlags().StringVarP(&AddCommitPFlag, "commit", "c", "", "The commit hash to check out. Overrides branch and tag.")
	viper.BindPFlag("commit", AddCmd.PersistentFlags().Lookup("commit"))

	AddCmd.PersistentFlags().BoolVarP(&AddSubmodulePFlag, "submodule", "s", false, "Add as a submodule. Git must be init&apos;d already.")
	viper.BindPFlag("submodule", AddCmd.PersistentFlags().Lookup("submodule"))

}

var AddCmd = &cobra.Command{

	Use: "add [sub-command] [git-repo-url]",

	Short: "Add a design, dsl, or generator to a project.",

	Long: AddLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In addCmd", "args", args)
		// Argument Parsing
		// [0]name:   url
		//     help:   The url of a git repository. May be any of the remote types (git@, http(s)).
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'url'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var url string

		if 0 < len(args) {

			url = args[0]
		}

		// [1]name:   location
		//     help:   The location for the design. Defaults to the first design path listed in the geb.yaml file.
		//     req'd:

		var location string

		if 1 < len(args) {

			location = args[1]
		}

		// HOFSTADTER_START cmd_run
		fmt.Println("geb add:",
			url,

			location,
		)
		opts := add.Options{
			Url:      url,
			Location: filepath.Join("vendor", location),
			Branch:   viper.GetString("branch"),
			Tag:      viper.GetString("tag"),
			Commit:   viper.GetString("commit"),
			Global:   viper.GetBool("global"),
		}

		err := add.AddGitRepo(opts)
		if err != nil {
			fmt.Printf("Error:\n%v\n", err)
			os.Exit(1)
		}
		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}

func init() {
	// add sub-commands to this command when present

}
