package system

// package commands

import (
	// HOFSTADTER_START import
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   init
// Usage:  init
// Parent: system

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var InitLong = `Intializes the geb tool and the ~/.geb dot folder.`

var InitCmd = &cobra.Command{

	Use: "init",

	Aliases: []string{
		"initialize",
		"setup",
	},

	Short: "Initialize the geb tool and files.",

	Long: InitLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In initCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		fmt.Println("Initializing geb and libraries")

		// get user and home dir
		u, err := user.Current()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		home := u.HomeDir

		const gebGoDirEnv = "$GOPATH/src/github.ibm.com/hofstadter-io/geb"
		gebGoDir := os.ExpandEnv(gebGoDirEnv)

		const dslGoDirEnv = "$GOPATH/src/github.ibm.com/hofstadter-io/dsl-library"
		dslGoDir := os.ExpandEnv(dslGoDirEnv)

		// possibly create dotfolder
		dotFolder := filepath.Join(home, ".geb")
		err = os.MkdirAll(dotFolder, 0755)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// copy in geb.yaml
		gebFileSrc := filepath.Join(gebGoDir, "lib/dotfolder/geb.yaml")
		gebFileDest := filepath.Join(dotFolder, "geb.yaml")
		fmt.Println("copy: ", gebFileSrc, gebFileDest)
		err = utils.CopyFile(gebFileSrc, gebFileDest)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// copy dsl folder
		dslFolderSrc := filepath.Join(dslGoDir)
		dslFolderDest := filepath.Join(dotFolder, "dsl")

		// possibly remove existing dir, (returns nil if non-existent)
		err = os.RemoveAll(dslFolderDest)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// copy in the dsl folder
		fmt.Println("copy: ", dslFolderSrc, dslFolderDest)
		err = utils.CopyFolder(dslFolderSrc, dslFolderDest)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
