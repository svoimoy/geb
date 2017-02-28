{{#with . as |Cmd| }}
{{#if Cmd.pflags}}
var (
{{#each Cmd.pflags}}	{{camel name }}PFlag {{type}}
	{{/each}}
)
{{/if}}

{{#if Cmd.flags}}
var (
{{#each Cmd.flags}}	{{camel name }}Flag {{type}}
	{{/each}}
)
{{/if}}
{{/with}}

