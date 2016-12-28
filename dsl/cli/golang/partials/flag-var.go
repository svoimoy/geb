{{#with . as |Cmd| }}
{{#if Cmd.pflags}}
var (
{{#each Cmd.pflags}}	{{ name }}PFlag {{type}}
	{{/each}}
)
{{/if}}

{{#if Cmd.flags}}
var (
{{#each Cmd.flags}}	{{ name }}Flag {{type}}
	{{/each}}
)
{{/if}}
/*
{{ Cmd }}
*/
{{/with}}

