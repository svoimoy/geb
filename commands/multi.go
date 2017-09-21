package commands

import (
	// HOFSTADTER_START import
	"fmt"
	"os"

	"github.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   multi
// Usage:  multi
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var MultiLong = `Generate stuff from data and template directories.`

var (
	MultiInputFlag []string

	MultiTemplateDirFlag []string

	MultiRepeatFlag []string

	MultiOutputDirFlag string
)

func init() {
	MultiCmd.Flags().StringArrayVarP(&MultiInputFlag, "input", "I", []string{}, "Paths to input files and directories, can be specified multiple times and override/extend.")
	viper.BindPFlag("input", MultiCmd.Flags().Lookup("input"))

	MultiCmd.Flags().StringArrayVarP(&MultiTemplateDirFlag, "template-dir", "T", []string{}, "Paths to template directories, can be specified multiple times and override/extend.")
	viper.BindPFlag("template-dir", MultiCmd.Flags().Lookup("template-dir"))

	MultiCmd.Flags().StringArrayVarP(&MultiRepeatFlag, "repeat", "R", []string{}, "Templates to repeat, form is &quot;.some.dotpath:template/path/{{dotpath.for.filename}}.ext&quot;")
	viper.BindPFlag("repeat", MultiCmd.Flags().Lookup("repeat"))

	MultiCmd.Flags().StringVarP(&MultiOutputDirFlag, "output-dir", "O", ".", "Path to the output directory which will prefix the template file/dir structure.")
	viper.BindPFlag("output-dir", MultiCmd.Flags().Lookup("output-dir"))

}

var MultiCmd = &cobra.Command{

	Use: "multi",

	Short: "Generate stuff from data and template directories.",

	Long: MultiLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In multiCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run
		fmt.Println("geb multi:")

		if len(MultiInputFlag) == 0 {
			fmt.Println("No inputs specified, use the '-I' flag.")
			os.Exit(1)
		}

		if len(MultiTemplateDirFlag) == 0 {
			fmt.Println("No templates specified, use the '-T' flag.")
			os.Exit(1)
		}

		if MultiOutputDirFlag == "" {
			fmt.Println("No output directory specified, use the '-O' flag.")
			os.Exit(1)
		}

		err := engine.GenerateMulti(
			MultiInputFlag,
			MultiTemplateDirFlag,
			MultiRepeatFlag,
			MultiOutputDirFlag,
		)

		if err != nil {
			fmt.Println("Error:\n", err)
			os.Exit(1)
		}

		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(MultiCmd)
}

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
