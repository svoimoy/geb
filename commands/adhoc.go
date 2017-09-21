package commands

import (
	// HOFSTADTER_START import
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/clbanning/mxj"
	"github.com/ghodss/yaml"
	"github.com/naoina/toml"

	"github.com/hofstadter-io/data-utils/io"
	"github.com/hofstadter-io/geb/engine"
	"github.com/hofstadter-io/hof-lang/lib/ast"
	"github.com/hofstadter-io/hof-lang/lib/parser"
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

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
	AdhocInputFlag string

	AdhocInputTypeFlag string

	AdhocFieldFlag string

	AdhocFlattenFlag int

	AdhocMultiFlag bool

	AdhocTemplateStringFlag string

	AdhocTemplateFileFlag string

	AdhocOutputFlag string

	AdhocOutputTypeFlag string
)

func init() {
	AdhocCmd.Flags().StringVarP(&AdhocInputFlag, "input", "I", "stdin", "path to an input file or directory, merged with the input command&apos;s arguments.")
	viper.BindPFlag("input", AdhocCmd.Flags().Lookup("input"))

	AdhocCmd.Flags().StringVarP(&AdhocInputTypeFlag, "input-type", "i", "auto", "input type, one of [yaml,json,toml,xml]")
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

	AdhocCmd.Flags().StringVarP(&AdhocOutputFlag, "output", "O", "stdout", "path to an output file or directory")
	viper.BindPFlag("output", AdhocCmd.Flags().Lookup("output"))

	AdhocCmd.Flags().StringVarP(&AdhocOutputTypeFlag, "output-type", "o", "", "output type from [yaml,json,toml,xml] (default &quot;json&quot;)")
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
		var inputContentType string
		var err error

		if AdhocInputTypeFlag == "auto" {
			if AdhocInputFlag == "stdin" {
				ctype, cerr := io.ReadAll(os.Stdin, &inputData)
				errExit(cerr)
				inputContentType = ctype
			} else {
				ctype, cerr := io.ReadFile(AdhocInputFlag, &inputData)
				errExit(cerr)
				inputContentType = ctype
			}
		} else {
			inputContentType = AdhocInputTypeFlag
			var content []byte

			if AdhocInputFlag == "stdin" {
				content, err = ioutil.ReadAll(os.Stdin)
				errExit(err)
			} else {
				content, err = ioutil.ReadFile(AdhocInputFlag)
				errExit(err)
			}

			// unmarshal into interface{}
			switch inputContentType {
			case "json":
				err = json.Unmarshal(content, &inputData)
				errExit(err)
			case "yaml", "yml":
				err = yaml.Unmarshal(content, &inputData)
				errExit(err)
			case "xml":
				mv, merr := mxj.NewMapXml(content)
				errExit(merr)
				inputData = map[string]interface{}(mv)
			case "toml":
				err = toml.Unmarshal(content, &inputData)
				errExit(err)
			case "hof":
				result, err := parser.ParseReader("", bytes.NewReader(content))
				errExit(err)

				hofFile := result.(ast.HofFile)
				hofData, err := hofFile.ToData()
				errExit(err)

				inputData = hofData
			default:
				fmt.Println("unknown input type: ", AdhocInputTypeFlag)
				os.Exit(1)
			}
		}

		// read in the template
		data := []byte("{{{json .}}}")
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
			case "json":
				data = []byte("{{{json .}}}")
			case "yaml", "yml":
				data = []byte("{{{yaml .}}}")
			case "xml":
				data = []byte("{{{xml .}}}")
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

func init() {
	// add sub-commands to this command when present

}

// HOFSTADTER_BELOW
