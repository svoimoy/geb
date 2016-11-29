{{#with dsl.cli as |CLI| }}
package cmd

import (
	{{#unless CLI.omit-root-run}}
	"fmt"
	{{/unless}}

	"github.com/spf13/cobra"
)

var (
{{#each CLI.flags}}	{{ name }}Flag string
	{{/each}}
)

func init() {
	//	cobra.OnInitialize(initConfig)
{{#each CLI.flags}}	RootCmd.PersistentFlags().StringVarP(&{{ name }}Flag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
{{/each}}
}

var (
	RootCmd = &cobra.Command{
		Use:   "{{ CLI.name }}",
		Short: "{{ CLI.short }}",
		Long:  `{{ CLI.long }}`,
		{{#unless CLI.omit-root-run}}
		Run: func(cmd *cobra.Command, args []string) {
			// HOFSTADTER_START root_cmd_func
			// Do Stuff Here
			fmt.Println("dostuff")
			// HOFSTADTER_END   root_cmd_func
		},
		{{/unless}}
	}
)
{{/with}}
