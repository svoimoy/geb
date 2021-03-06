package commands

import (
	// HOFSTADTER_START import
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports
	// infered imports

	"github.com/spf13/cobra"
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var GebLong = `Hofstadter is a Framework
for building data-centric
Platforms. geb is the tool.
`

var (
	RootCmd = &cobra.Command{

		Use: "geb",

		Short: "geb is the Hofstadter framework CLI tool",

		Long: GebLong,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Argument Parsing

			// HOFSTADTER_START cmd_persistent_prerun

			// short circuit for system child commands
			if cmd.Parent().Name() == "system" {
				return
			} else {
				// temporary short circuit
				return
			}

			// Look for .geb dotfolder

			// get user and home dir
			u, err := user.Current()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			home := u.HomeDir

			dotfolder := filepath.Join(home, ".geb")
			gebFile := filepath.Join(home, ".geb/geb.yaml")

			_, err = os.Lstat(dotfolder)
			if err != nil {
				if _, ok := err.(*os.PathError); ok {
					fmt.Println("The geb dotfolder  ($HOME/.geb/) appears to be missing :(\n\n  Try running 'geb sys init'.\n")
					os.Exit(1)
				} else {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			}

			_, err = os.Lstat(gebFile)
			if err != nil {
				if _, ok := err.(*os.PathError); ok {
					fmt.Println("The geb system config file ($HOME/.geb/geb.yaml) appears to be missing :(\n\n  Try running 'geb sys init'.\n")
					os.Exit(1)
				} else {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			}

			// look for gopath
			if len(os.ExpandEnv("${GOPATH}")) == 0 {
				fmt.Println("GOPATH environment variable not set.\n\n  Please see https://github.com/golang/go/wiki/GOPATH\n")
				os.Exit(1)
			}

			// HOFSTADTER_END   cmd_persistent_prerun
		},
	}
)

// HOFSTADTER_BELOW
