package commands

// hof  -  hof

import (
	// HOFSTADTER_START import
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hofstadter-io/data-utils/io"
	"github.com/hofstadter-io/hof-lang/lib/ast"
	"github.com/hofstadter-io/hof-lang/lib/parser"

	"encoding/json"
	"github.com/clbanning/mxj"
	"github.com/ghodss/yaml"
	"github.com/naoina/toml"
	// HOFSTADTER_END   import

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var (
	RootInputFlag      string
	RootInputTypeFlag  string
	RootOutputFlag     string
	RootOutputTypeFlag string
)

func init() {
	RootCmd.Flags().StringVarP(&RootInputFlag, "input", "I", "stdin", "path to an input file or directory, merged with the input command&apos;s arguments.")
	viper.BindPFlag("input", RootCmd.Flags().Lookup("input"))

	RootCmd.Flags().StringVarP(&RootInputTypeFlag, "input-type", "i", "auto", "input type, one of [yaml,json,toml,xml,hof]")
	viper.BindPFlag("input-type", RootCmd.Flags().Lookup("input-type"))

	RootCmd.Flags().StringVarP(&RootOutputFlag, "output", "O", "stdout", "path to an output file or directory")
	viper.BindPFlag("output", RootCmd.Flags().Lookup("output"))

	RootCmd.Flags().StringVarP(&RootOutputTypeFlag, "output-type", "o", "yaml", "output type from [yaml,hof,json]")
	viper.BindPFlag("output-type", RootCmd.Flags().Lookup("output-type"))

}

var (
	RootCmd = &cobra.Command{

		Use: "hof",

		Short: "The Hofstadter transpiler tool.",

		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("In hofCmd", "args", args)
			// Argument Parsing

			// HOFSTADTER_START cmd_run
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

			if RootInputTypeFlag == "auto" {
				if RootInputFlag == "stdin" {
					ctype, cerr := io.ReadAll(os.Stdin, &inputData)
					errExit(cerr)
					inputContentType = ctype
				} else {
					ctype, cerr := io.ReadFile(RootInputFlag, &inputData)
					errExit(cerr)
					inputContentType = ctype
				}
			} else {
				inputContentType = RootInputTypeFlag
				var content []byte

				if RootInputFlag == "stdin" {
					content, err = ioutil.ReadAll(os.Stdin)
					errExit(err)
				} else {
					content, err = ioutil.ReadFile(RootInputFlag)
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
					fmt.Println("unknown input type: ", RootInputTypeFlag)
					os.Exit(1)
				}
			}

			output := ""
			switch RootOutputTypeFlag {

			case "json":
				bytes, err := json.MarshalIndent(inputData, "", "  ")
				errExit(err)
				output = string(bytes)

			case "yaml", "yml":
				bytes, err := yaml.Marshal(inputData)
				errExit(err)
				output = string(bytes)

			case "hof":
				var tree ast.HofFile
				ret, err := tree.FromData(inputData)
				/*
				fmt.Println("============================")
				fmt.Printf("%# v\n", pretty.Formatter(ret))
				fmt.Println("============================")
				*/
				errExit(err)
				data, err := ret.String("")
				output = data

			default:
				fmt.Println("unknown output type: ", RootOutputTypeFlag)
				os.Exit(1)
			}

			if RootOutputFlag == "stdout" {
				output = "\n\n" + output
				fmt.Println(output)
			} else {
				err := ioutil.WriteFile(RootOutputFlag, []byte(output), 0644)
				errExit(err)
			}

			// HOFSTADTER_END   cmd_run
		},
	}
)

// HOFSTADTER_BELOW
