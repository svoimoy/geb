func New{{camelT TYP.name}}() *{{camelT TYP.name}} {
	return &{{camelT TYP.name}}{
		{{#each TYP.fields as |F|}}
		{{#if (contains F.type "array") ~}}
		{{camelT F.name}}: {{>types/golang/type.go F.type ~}}{},
		{{else if (contains F.type "map") ~}}
		{{camelT F.name}}: {{>types/golang/type.go F.type ~}}{},
		{{else if (contains F.type "*") ~}}
		{{camelT F.name}}: New{{>types/golang/type-def-no-pkg.go F.type ~}}(),
		{{/if ~}}
		{{/each}}
	}
	// loop over fields looking for pointers
}

/*
{{{yaml TYP}}}
*/

