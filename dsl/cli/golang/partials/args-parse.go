// Argument Parsing
{{#with . as |Cmd| }}
{{#each Cmd.args}}
{{#with . as |arg|}}
// [{{@index}}]name:   {{arg.name}}
//     help:   {{{arg.help}}}
//     req'd:  {{arg.required}}
	{{#if arg.required}}
	if {{@index}} >= len(args) {
		fmt.Printf("missing required argument: '{{arg.name}}'")
		cmd.Usage()
		os.Exit(1)
	}
	{{/if}}

	var {{camel arg.name}} {{> go-type.go arg.type}}
	{{#if arg.default}}
		{{#if (eq arg.type "string")}}
		{{camel arg.name}} = "{{arg.default}}"
		{{else}}
		{{camel arg.name}} = {{arg.default}}
		{{/if}}
	{{/if}}

	if {{@index}} < len(args) {
	{{#if arg.rest}}
		{{#if (eq arg.type "array:string")}}
			{{camel arg.name}} = args[{{@index}}:]
		{{else}}
			// 'rest' args can only be of type 'array:string' currently...
			// perhaps a loop of parsing
		{{/if}}
	{{else if (eq arg.type "string")}}
			{{camel arg.name}} = args[{{@index}}]
	{{else}}
			{{camel arg.name}}Arg := args[{{@index}}]
			var err error
			{{> common/golang/parse/builtin.go IN_NAME=(concat2 (camel arg.name) "Arg") OUT_NAME=(camel arg.name) TYP=arg.type}}
			if err != nil {
				fmt.Printf("argument of wrong type. expected: '{{arg.type}}' got error: %v", err)
				cmd.Usage()
				os.Exit(1)
			}
	{{/if}}
	}
	
	

{{/with}}
{{/each}}
{{/with}}

