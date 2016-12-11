// Argument Parsing
{{#with . as |Cmd| }}
{{#each Cmd.args}}
{{#with . as |arg|}}
// [{{@index}}]name:   {{arg.name}}
//     help:   {{{arg.help}}}
//     req'd:  {{arg.required}}
{{#if arg.required}}
if {{@index}} >= len(args) {
	fmt.Println("not enough args supplied")
	return
}
{{/if}}
{{#if arg.rest}}
var {{arg.name}} {{> go-type.go arg.type}}
if {{@index}} < len(args) {
	{{arg.name}} = args[{{@index}}:]
}
{{else}}
var {{arg.name}} {{arg.type}}
if {{@index}} < len(args) {
	{{arg.name}} = args[{{@index}}]
}
{{/if}}

fmt.Println("arg[{{@index}}] = ", {{arg.name}})

{{/with}}
{{/each}}
{{/with}}

