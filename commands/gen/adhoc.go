package gen

// package subcommands

import (
	// HOFSTADTER_START import
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
	"github.com/naoina/toml"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   adhoc
// Usage:  adhoc
// Parent: gen

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var AdhocLong = `Generate something from data and a template.`

var (
	inputFlag          string
	inputTypeFlag      string
	fieldFlag          string
	flattenFlag        int
	templateStringFlag string
	templateFileFlag   string
	outputFlag         string
)

func init() {
	AdhocCmd.Flags().StringVarP(&inputFlag, "input", "i", "stdin", "path to an input file or directory")
	viper.BindPFlag("input", AdhocCmd.Flags().Lookup("input"))

	AdhocCmd.Flags().StringVarP(&inputTypeFlag, "input-type", "t", "yaml", "type of the data in the input file or directory")
	viper.BindPFlag("input-type", AdhocCmd.Flags().Lookup("input-type"))

	AdhocCmd.Flags().StringVarP(&fieldFlag, "field", "f", ".", "a dotpath into the data to be used for rendering")
	viper.BindPFlag("field", AdhocCmd.Flags().Lookup("field"))

	AdhocCmd.Flags().IntVarP(&flattenFlag, "flatten", "", 0, "flattend nested arrays by N levels")
	viper.BindPFlag("flatten", AdhocCmd.Flags().Lookup("flatten"))

	AdhocCmd.Flags().StringVarP(&templateStringFlag, "template-string", "T", "{{{yaml .}}}", "Template contents to render with.")
	viper.BindPFlag("template-string", AdhocCmd.Flags().Lookup("template-string"))

	AdhocCmd.Flags().StringVarP(&templateFileFlag, "template-file", "F", "", "Path to the template file.")
	viper.BindPFlag("template-file", AdhocCmd.Flags().Lookup("template-file"))

	AdhocCmd.Flags().StringVarP(&outputFlag, "output", "o", "stdout", "path to an output file or directory")
	viper.BindPFlag("output", AdhocCmd.Flags().Lookup("output"))

}

var AdhocCmd = &cobra.Command{

	Use: "adhoc",

	Aliases: []string{
		"on-the-fly",
	},

	Short: "Generate something from data and a template.",

	Long: AdhocLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In adhocCmd", "args", args)
		// Argument Parsing

		// HOFSTADTER_START cmd_run

		// to shorten the code below
		errExit := func(err error) {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// read in data
		var inputData map[string]interface{}
		var data []byte
		var err error
		if inputFlag == "stdin" {
			data, err = ioutil.ReadAll(os.Stdin)
			errExit(err)
		} else {
			data, err = ioutil.ReadFile(inputFlag)
			errExit(err)
		}

		// unmarshal into interface{}
		switch inputTypeFlag {
		case "yaml", "yml":
			err = yaml.Unmarshal(data, &inputData)
			errExit(err)
		case "json":
			err = json.Unmarshal(data, &inputData)
			errExit(err)
		case "toml":
			err = toml.Unmarshal(data, &inputData)
			errExit(err)
		default:
			fmt.Println("unknown input type: ", inputTypeFlag)
			os.Exit(1)
		}

		// read in the template
		data = []byte(templateStringFlag)
		if templateFileFlag != "" {
			data, err = ioutil.ReadFile(templateFileFlag)
			errExit(err)
		}

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
