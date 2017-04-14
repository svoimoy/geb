package gen

// package subcommands

import (
	// HOFSTADTER_START import
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   adhoc
// Usage:  adhoc <templateFile>
// Parent: gen

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var AdhocLong = `Generate something from data and a template.`

var (
	inputFlag     string
	inputTypeFlag string
	fieldFlag     string
	outputFlag    string
)

func init() {
	AdhocCmd.Flags().StringVarP(&inputFlag, "input", "i", "stdin", "path to an input file or directory")
	viper.BindPFlag("input", AdhocCmd.Flags().Lookup("input"))

	AdhocCmd.Flags().StringVarP(&inputTypeFlag, "input-type", "t", "yaml", "type of the data in the input file or directory")
	viper.BindPFlag("input-type", AdhocCmd.Flags().Lookup("input-type"))

	AdhocCmd.Flags().StringVarP(&fieldFlag, "field", "f", ".", "a dotpath into the data to be used for rendering")
	viper.BindPFlag("field", AdhocCmd.Flags().Lookup("field"))

	AdhocCmd.Flags().StringVarP(&outputFlag, "output", "o", "stdout", "path to an output file or directory")
	viper.BindPFlag("output", AdhocCmd.Flags().Lookup("output"))

}

var AdhocCmd = &cobra.Command{

	Use: "adhoc <templateFile>",

	Aliases: []string{
		"on-the-fly",
	},

	Short: "Generate something from data and a template.",

	Long: AdhocLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In adhocCmd", "args", args)
		// Argument Parsing
		// [0]name:   template-file
		//     help:   Path to the template file.
		//     req'd:  true
		if 0 >= len(args) {
			fmt.Println("missing required argument: 'template-file'\n")
			cmd.Usage()
			os.Exit(1)
		}

		var templateFile string

		if 0 < len(args) {

			templateFile = args[0]
		}

		// HOFSTADTER_START cmd_run

		// to shorten the code below
		errExit := func(err error) {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// read in data
		var inputData interface{}
		var data []byte
		var err error
		if inputFlag == "stdin" {
			data, err = ioutil.ReadAll(os.Stdin)
			errExit(err)
		} else {
			data, err = ioutil.ReadFile(inputFlag)
			errExit(err)
		}

		// need to switch on input filename extension here
		// unmarshal into interface{}
		err = yaml.Unmarshal(data, &inputData)
		errExit(err)

		// read in the template
		data, err = ioutil.ReadFile(templateFile)
		errExit(err)

		templateData := string(data)

		// generate
		outputData, err := engine.GenerateAdhoc(inputData, fieldFlag, templateData)
		errExit(err)

		// write the output
		if outputFlag == "stdout" {
			fmt.Println(outputData)
		} else {
			err := ioutil.WriteFile(outputFlag, []byte(outputData), 0644)
			errExit(err)
		}

		// HOFSTADTER_END   cmd_run
	},
}

// HOFSTADTER_BELOW
