{{#with . as |Cmd| }}
{{#if Cmd.pflags}}
func init() {
{{#each Cmd.pflags}}
{{#if Cmd.parent}}
	{{#if (eq type "int")}}
	{{Cmd.name}}Cmd.PersistentFlags().IntVarP(&{{ name }}PFlag, "{{long}}", "{{short}}", {{default}}, "{{help}}")
	{{else}}
	{{Cmd.name}}Cmd.PersistentFlags().StringVarP(&{{ name }}PFlag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	{{/if}}
	viper.BindPFlag("{{long}}", {{Cmd.name}}Cmd.PersistentFlags().Lookup("{{long}}"))
{{else}}
	{{#if (eq type "int")}}
	RootCmd.PersistentFlags().IntVarP(&{{ name }}PFlag, "{{long}}", "{{short}}", {{default}}, "{{help}}")
	{{else}}
	RootCmd.PersistentFlags().StringVarP(&{{ name }}PFlag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	{{/if}}
	viper.BindPFlag("{{long}}", RootCmd.PersistentFlags().Lookup("{{long}}"))
{{/if}}

{{/each}}
}
{{/if}}

{{#if Cmd.flags}}
func init() {
{{#each Cmd.flags}}
{{#if Cmd.parent}}
	{{#if (eq type "int")}}
	{{Cmd.name}}Cmd.Flags().IntVarP(&{{ name }}Flag, "{{long}}", "{{short}}", {{default}}, "{{help}}")
	{{else}}
	{{Cmd.name}}Cmd.Flags().StringVarP(&{{ name }}Flag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	{{/if}}
	viper.BindPFlag("{{long}}", {{Cmd.name}}Cmd.Flags().Lookup("{{long}}"))
{{else}}
	{{#if (eq type "int")}}
	RootCmd.Flags().IntVarP(&{{ name }}Flag, "{{long}}", "{{short}}", {{default}}, "{{help}}")
	{{else}}
	RootCmd.Flags().StringVarP(&{{ name }}Flag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	{{/if}}
	viper.BindPFlag("{{long}}", RootCmd.Flags().Lookup("{{long}}"))
{{/if}}

{{/each}}
}
{{/if}}
{{/with}}
