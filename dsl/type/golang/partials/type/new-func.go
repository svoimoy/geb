func New{{camelT TYP.name}}() *{{camelT TYP.name}} {
	return &{{camelT TYP.name}}{
		{{#each TYP.fields as |F|}}
		{{#if (contains F.type "array") ~}}
		{{camelT F.name}}: {{>type/golang/type.go F ~}}{},
		{{else if (contains F.type "map") ~}}
		{{camelT F.name}}: {{>type/golang/type.go F ~}}{},
		{{else if (contains F.type "*") ~}}
		{{#with (getsuffix (getsuffix F.type ":") "*") as |T| ~}}
			{{#if (builtin T)}}
			{{camelT F.name}}: new({{ T}}),
			{{else if (contains T "/") ~}}
			{{camelT F.name}}: new({{> type/golang/package.go TYP=F ~}}{{>type/golang/type-def-no-pkg.go F ~}}),
			{{else}}{{! we should check if it is defined in the known data type}}
			{{camelT F.name}}: {{> type/golang/package.go TYP=F ~}}New{{>type/golang/type-def-no-pkg.go F ~}}(),
			{{/if}}
		{{/with}}
		{{/if ~}}
		{{/each}}
	}
	// loop over fields looking for pointers
}

