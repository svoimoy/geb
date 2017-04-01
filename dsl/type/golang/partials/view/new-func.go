func New{{camelT TYP.name}}View{{camelT VIEW.name}}() *{{camelT TYP.name}}View{{camelT VIEW.name}} {
	return &{{camelT TYP.name}}View{{camelT VIEW.name}}{
		{{#each VIEW.fields as |F|}}
			{{#if F.new-func}}
				{{camelT F.name}}: {{ F.new-func }},
			{{else if (contains F.type "array") ~}}
				{{camelT F.name}}: {{>type/golang/type.go TYP=F ~}}{},
			{{else if (contains F.type "map") ~}}
				{{camelT F.name}}: {{>type/golang/type.go TYP=F ~}}{},
			{{else if (contains F.type "channel") ~}}
				{{camelT F.name}}: {{>type/golang/type.go TYP=F ~}}{},
			{{else if (contains F.type "*") ~}}
				{{#with (getsuffix (getsuffix F.type ":") "*") as |T| ~}}
					{{#if (builtin T)}}
					{{camelT F.name}}: new({{ T}}),
					{{else if (contains T "/") ~}}
					{{camelT F.name}}: new({{> type/golang/package.go TYP=T ~}}{{>type/golang/type-def-no-pkg.go T ~}}),
					{{else}}{{! we should check if it is defined in the known data type}}
					{{camelT F.name}}: {{> type/golang/package.go TYP=T ~}}New{{>type/golang/type-def-no-pkg.go T ~}}(),
					{{/if}}
				{{/with}}
			{{/if ~}}
		{{/each}}
	}
}

