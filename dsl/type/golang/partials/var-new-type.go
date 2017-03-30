{{#if (builtin TYP.type)}}
	// builtin
	var {{camel TYP.name}} {{TYP.type}}

{{else}}
	// user-defined
	var {{camel TYP.name }} {{!formatting block ~}}
	{{> type/golang/package.go TYP=TYP ~}}
	{{#with (getsuffix (getsuffix TYP.type ":") "*") as |T1|}}
	{{#with (trimto_last T1 "." false) as |T|}}
		{{#if (builtin T) ~}}
			{{ T ~}}
		{{else ~}}
			{{camelT T ~}}
		{{/if }}
	{{/with}}
	{{/with}}
{{/if}}
