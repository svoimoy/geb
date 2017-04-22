package system

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   dev-mode
// Usage:  dev-mode
// Parent: system

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var DevModeCmd = &cobra.Command{

	Use: "dev-mode",

	Aliases: []string{
		"dev",
	},

	Short: "Put geb in development mode.",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In dev-modeCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		fmt.Println("geb system dev-mode: Putting geb in developer mode.")

		// get user and home dir
		u, err := user.Current()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		home := u.HomeDir

		const gebGoDirEnv = "$GOPATH/src/github.com/hofstadter-io/geb"
		gebGoDir := os.ExpandEnv(gebGoDirEnv)

		const dslGoDirEnv = "$GOPATH/src/github.com/hofstadter-io/dsl-library"
		dslGoDir := os.ExpandEnv(dslGoDirEnv)

		// possibly create dotfolder
		dotFolder := filepath.Join(home, ".geb")
		err = os.MkdirAll(dotFolder, 0755)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// copy in geb.yaml
		gebFileSrc := filepath.Join(gebGoDir, "lib/dotfolder/dev.yaml")
		gebFileDest := filepath.Join(dotFolder, "geb.yaml")
		fmt.Println("copy: ", gebFileSrc, gebFileDest)
		err = utils.CopyFile(gebFileSrc, gebFileDest)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// link dsl folder
		dslFolderSrc := filepath.Join(dslGoDir)
		dslFolderDest := filepath.Join(dotFolder, "dsl")

		// possibly remove existing dir, (returns nil if non-existent)
		err = os.RemoveAll(dslFolderDest)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// link in the dsl folder
		fmt.Println("link: ", dslFolderSrc, dslFolderDest)
		err = os.Symlink(dslFolderSrc, dslFolderDest)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// echo path update
		fmt.Println("Add 'PATH=$PATH:$GOPATH/src/github.com/hofstadter-io/geb' to your .profile, .bashrc, or which ever you may use.")

		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
