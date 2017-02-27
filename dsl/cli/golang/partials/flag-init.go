{{#with . as |Cmd| }}
{{#if Cmd.pflags}}
func init() {
{{#each Cmd.pflags}}
{{#if Cmd.parent}}
	{{> cli/golang/flag-def.go FLAG=. CMDNAME=(camelT Cmd.name) PERSIST="Persitent"}}
{{else}}
	{{> cli/golang/flag-def.go FLAG=. CMDNAME="Root" PERSIST="Persitent"}}
{{/if}}

{{/each}}
}
{{/if}}

{{#if Cmd.flags}}
func init() {
{{#each Cmd.flags}}
{{#if Cmd.parent}}
	{{> cli/golang/flag-def.go FLAG=. CMDNAME=(camelT Cmd.name) PERSIST=""}}
{{else}}
	{{> cli/golang/flag-def.go FLAG=. CMDNAME="Root" PERSIST=""}}
{{/if}}

{{/each}}
}
{{/if}}
{{/with}}
