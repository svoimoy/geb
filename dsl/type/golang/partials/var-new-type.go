{{#if (builtin TYP.type)}}
	// builtin
	var {{camel TYP.name}} {{TYP.type}}

{{else}}
	// user-defined
	var {{camel TYP.name }} {{!formatting block ~}}
	{{> type/golang/type-def.go TYP=TYP ~}}
{{/if}}
