package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/geb/commands/add"
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
	AddGlobalPFlag boolean

	AddBranchPFlag string

	AddTagPFlag string

	AddCommitPFlag string
)

func init() {
	/* unknown Flag type in:
	ctx_path: dsl.cli.commands.[0].pflags.[0]
	default: false
	help: add the package to the global context in ~/.geb/dsl/... Ignored by designs
	long: global
	name: global
	parent: geb.add
	parent_path: dsl.cli.commands.[0]
	pkg_path: cli/commands/[0]/pflags
	pkgPath: geb/add/global
	short: g
	type: boolean

	*/
	viper.BindPFlag("global", AddCmd.PersistentFlags().Lookup("global"))

	AddCmd.PersistentFlags().StringVarP(&AddBranchPFlag, "branch", "b", "master", "The branch to check out.")
	viper.BindPFlag("branch", AddCmd.PersistentFlags().Lookup("branch"))

	AddCmd.PersistentFlags().StringVarP(&AddTagPFlag, "tag", "t", "", "The tag to check out. Overrides branch.")
	viper.BindPFlag("tag", AddCmd.PersistentFlags().Lookup("tag"))

	AddCmd.PersistentFlags().StringVarP(&AddCommitPFlag, "commit", "c", "", "The commit hash to check out. Overrides branch and tag.")
	viper.BindPFlag("commit", AddCmd.PersistentFlags().Lookup("commit"))

}

var AddCmd = &cobra.Command{

	Use: "add [sub-command] [git-repo-url]",

	Short: "Add a design, dsl, or generator to a project.",

	Long: AddLong,
}

func init() {
	RootCmd.AddCommand(AddCmd)
}

func init() {
	// add sub-commands to this command when present

	AddCmd.AddCommand(add.DesignCmd)
	AddCmd.AddCommand(add.DslCmd)
	AddCmd.AddCommand(add.GeneratorCmd)
}
