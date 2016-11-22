package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
{{#each CLI.flags}}	{{ name }}Flag string
	{{/each}}
)

func init() {
	//	cobra.OnInitialize(initConfig)
{{#each CLI.flags}}	RootCmd.PersistentFlags().StringVarP(&{{ name }}Flag, "{{name}}", "{{short}}", "{{default}}", "{{help}}")
{{/each}}
}

var (
	RootCmd = &cobra.Command{
		Use:   "{{ CLI.name }}",
		Short: "{{ CLI.short }}",
		Long:  `{{ CLI.long }}`,
		Run: func(cmd *cobra.Command, args []string) {
			// HOFSTADTER_START root_cmd_func
			// Do Stuff Here
			fmt.Println("dostuff")
			// HOFSTADTER_END   root_cmd_func
		},
	}
)
