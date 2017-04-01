// Argument Parsing
{{#with . as |Cmd| }}
{{#each Cmd.args}}
{{#with . as |arg|}}
// [{{@index}}]name:   {{arg.name}}
//     help:   {{{arg.help}}}
//     req'd:  {{arg.required}}
{{#if arg.required}}
if {{@index}} >= len(args) {
	cmd.Usage()
	return
}
{{/if}}
{{#if arg.rest}}
var {{camel arg.name}} {{> go-type.go arg.type}}
if {{@index}} < len(args) {
	{{camel arg.name}} = args[{{@index}}:]
}
{{else}}
var {{camel arg.name}} {{arg.type}}
if {{@index}} < len(args) {
	{{camel arg.name}} = args[{{@index}}]
}
{{/if}}

{{/with}}
{{/each}}
{{/with}}

