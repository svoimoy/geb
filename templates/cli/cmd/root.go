package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
{{#each dsl.cli.flags}}	{{ name }}Flag string
	{{/each}}
)

func init() {
	//	cobra.OnInitialize(initConfig)
{{#each dsl.cli.flags}}	RootCmd.PersistentFlags().StringVarP(&{{ name }}Flag, "{{name}}", "{{short}}", "{{default}}", "{{help}}")
{{/each}}
}

var (
	RootCmd = &cobra.Command{
		Use:   "{{ dsl.cli.name }}",
		Short: "{{ dsl.cli.short }}",
		Long:  `{{ dsl.cli.long }}`,
		Run: func(cmd *cobra.Command, args []string) {
			// HOFSTADTER_START root_cmd_func
			// Do Stuff Here
			fmt.Println("dostuff")
			// HOFSTADTER_END   root_cmd_func
		},
	}
)
