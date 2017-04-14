package commands

// package commands

import (
	// HOFSTADTER_START import
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/naoina/toml"
	// "gopkg.in/yaml.v2"

	"github.ibm.com/hofstadter-io/geb/engine"
	// HOFSTADTER_END   import

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   adhoc
// Usage:  adhoc
// Parent: geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var AdhocLong = `Generate something from data and a template.`

var (
	AdhocInputFlag          string
	AdhocInputTypeFlag      string
	AdhocFieldFlag          string
	AdhocFlattenFlag        int
	AdhocMultiFlag          bool
	AdhocTemplateStringFlag string
	AdhocTemplateFileFlag   string
	AdhocOutputFlag         string
	AdhocOutputTypeFlag     string
)

func init() {
	AdhocCmd.Flags().StringVarP(&AdhocInputFlag, "input", "i", "stdin", "path to an input file or directory")
	viper.BindPFlag("input", AdhocCmd.Flags().Lookup("input"))

	AdhocCmd.Flags().StringVarP(&AdhocInputTypeFlag, "input-type", "I", "yaml", "input type from [yaml,json,toml]")
	viper.BindPFlag("input-type", AdhocCmd.Flags().Lookup("input-type"))

	AdhocCmd.Flags().StringVarP(&AdhocFieldFlag, "field", "f", ".", "a dotpath into the data to be used for rendering")
	viper.BindPFlag("field", AdhocCmd.Flags().Lookup("field"))

	AdhocCmd.Flags().IntVarP(&AdhocFlattenFlag, "flatten", "", 0, "flattend nested arrays by N levels")
	viper.BindPFlag("flatten", AdhocCmd.Flags().Lookup("flatten"))

	AdhocCmd.Flags().BoolVarP(&AdhocMultiFlag, "multi", "", false, "the output is an array and each element should be put through the template. In this case the output flag should also be specified with a template for determining the path/to/file.out")
	viper.BindPFlag("multi", AdhocCmd.Flags().Lookup("multi"))

	AdhocCmd.Flags().StringVarP(&AdhocTemplateStringFlag, "template-string", "T", "", "Template contents to render with. Default: &apos;{{{&lt;output-type&gt; .}}}&apos;")
	viper.BindPFlag("template-string", AdhocCmd.Flags().Lookup("template-string"))

	AdhocCmd.Flags().StringVarP(&AdhocTemplateFileFlag, "template-file", "t", "", "Path to the template file.")
	viper.BindPFlag("template-file", AdhocCmd.Flags().Lookup("template-file"))

	AdhocCmd.Flags().StringVarP(&AdhocOutputFlag, "output", "o", "stdout", "path to an output file or directory")
	viper.BindPFlag("output", AdhocCmd.Flags().Lookup("output"))

	AdhocCmd.Flags().StringVarP(&AdhocOutputTypeFlag, "output-type", "O", "", "output type from [yaml,json,toml]")
	viper.BindPFlag("output-type", AdhocCmd.Flags().Lookup("output-type"))

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
		var inputData interface{}
		var data []byte
		var err error
		if AdhocInputFlag == "stdin" {
			data, err = ioutil.ReadAll(os.Stdin)
			errExit(err)
		} else {
			data, err = ioutil.ReadFile(AdhocInputFlag)
			errExit(err)
		}

		// unmarshal into interface{}
		switch AdhocInputTypeFlag {
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
			fmt.Println("unknown input type: ", AdhocInputTypeFlag)
			os.Exit(1)
		}

		// read in the template
		data = []byte("{{{yaml .}}}")
		tsF := AdhocTemplateStringFlag != ""
		tfF := AdhocTemplateFileFlag != ""
		otF := AdhocOutputTypeFlag != ""

		if (tsF && tfF) || (tsF && otF) || (tfF && otF) {
			fmt.Println("cannot specify only one of template-string, template-file, or output-type flags")
			fmt.Printf("%q %q %q\n", AdhocTemplateStringFlag, AdhocTemplateFileFlag, AdhocOutputTypeFlag)

			os.Exit(1)
		} else if tsF {
			data = []byte(AdhocTemplateStringFlag)
		} else if tfF {
			data, err = ioutil.ReadFile(AdhocTemplateFileFlag)
			errExit(err)
		} else if otF {
			switch AdhocOutputTypeFlag {
			case "yaml", "yml":
				data = []byte("{{{yaml .}}}")
			case "json":
				data = []byte("{{{json .}}}")
			case "toml":
				data = []byte("{{{toml .}}}")
			default:
				fmt.Printf("unknown output-type: %q\n", AdhocOutputTypeFlag)
				os.Exit(1)
			}
		}
		templateData := string(data)

		// generate
		outputData, err := engine.GenerateAdhoc(inputData, AdhocFieldFlag, templateData)
		errExit(err)

		// write the output
		if AdhocOutputFlag == "stdout" {
			fmt.Println(outputData)
		} else {
			err := ioutil.WriteFile(AdhocOutputFlag, []byte(outputData), 0644)
			errExit(err)
		}

		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(AdhocCmd)
}

// HOFSTADTER_BELOW
