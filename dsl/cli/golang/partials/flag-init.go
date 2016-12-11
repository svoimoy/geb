{{#with . as |Cmd| }}
{{#if Cmd.pflags}}
func init() {
{{#each Cmd.pflags}}
{{#if Cmd.parent}}
	{{Cmd.name}}Cmd.PersistentFlags().StringVarP(&{{ name }}PFlag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	viper.BindPFlag("{{long}}", {{Cmd.name}}Cmd.PersistentFlags().Lookup("{{long}}"))
{{else}}
	RootCmd.PersistentFlags().StringVarP(&{{ name }}PFlag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	viper.BindPFlag("{{long}}", RootCmd.PersistentFlags().Lookup("{{long}}"))
{{/if}}

{{/each}}
}
{{/if}}

{{#if Cmd.flags}}
func init() {
{{#each Cmd.flags}}
{{#if Cmd.parent}}
	{{Cmd.name}}Cmd.Flags().StringVarP(&{{ name }}Flag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	viper.BindPFlag("{{long}}", {{Cmd.name}}Cmd.Flags().Lookup("{{long}}"))
{{else}}
	RootCmd.Flags().StringVarP(&{{ name }}Flag, "{{long}}", "{{short}}", "{{default}}", "{{help}}")
	viper.BindPFlag("{{long}}", RootCmd.Flags().Lookup("{{long}}"))
{{/if}}

{{/each}}
}
{{/if}}
{{/with}}
