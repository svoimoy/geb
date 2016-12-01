{{#with . as |Cmd| }}
{{#if Cmd.pflags}}
var (
{{#each Cmd.pflags}}	{{ name }}PFlag {{#if type}}{{type}}{{else}}string{{/if}}
	{{/each}}
)
{{/if}}

{{#if Cmd.flags}}
var (
{{#each Cmd.flags}}	{{ name }}Flag {{#if type}}{{type}}{{else}}string{{/if}}
	{{/each}}
)
{{/if}}
{{/with}}
