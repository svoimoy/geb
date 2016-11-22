package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aymerick/raymond"
	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	configFile  string
	designDir   string
	templateDir string
	outputDir   string
	generators  string
)

func init() {
	//	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "geb.yaml", "geb config file for your projectd.")
	RootCmd.PersistentFlags().StringVarP(&designDir, "design-dir", "d", "design", "the design files directory. (default ./design)")
	RootCmd.PersistentFlags().StringVarP(&templateDir, "template-paths", "t", "main.go", "base templates directory. (default ./templates:~/.hofstadter/templates)")
	RootCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", "output", "the output files directory. (default ./output)")
	RootCmd.PersistentFlags().StringVarP(&generators, "generators", "g", "all", "which generator to run. (defaults to all found)")

	/*
		viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
		viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
		viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
		viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
		viper.SetDefault("license", "apache")
	*/
}

var (
	RootCmd = &cobra.Command{
		Use:   "geb",
		Short: "geb is a data centric code generator",
		Long: `A data centric code generator which
combines yaml and handlebar templates
to genereate all of the codes.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			fmt.Println(configFile, designDir, templateDir, outputDir, generators)
			dostuff()
		},
	}
)

func dostuff() {

	// fmt.Println("hof - data + templates = profit")

	raw_template, err := ioutil.ReadFile("templates/cli/cmd/root.go")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	source := string(raw_template)
	// parse template
	tpl, err := raymond.Parse(source)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	raw_data, err := ioutil.ReadFile("design/cli.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(raw_data), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	result, err := tpl.Exec(data)
	if err != nil {
		panic(err)
	}

	/*
	   t := template.Must(template.New("test-template").Parse(temp))
	   s := t.Execute(os.Stdout, data)
	*/
	fmt.Printf("%s\n", result)
}
