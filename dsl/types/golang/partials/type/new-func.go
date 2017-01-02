func New{{camelT TYP.name}}() *{{camelT TYP.name}} {
	return &{{camelT TYP.name}}{
		{{#each TYP.fields as |F|}}
		{{#if (contains F.type "array") ~}}
		{{> types/golang/type/new-func-field.go NAME=F.name TYP=F MOD=(ternary (trimsuffix F.type (trimfrom F.type "*" true)) (trimsuffix F.type (trimfrom F.type ":" true))) }}
		{{else if (contains F.type "map") ~}}
		{{> types/golang/type/new-func-field.go NAME=F.name TYP=F MOD=(ternary (trimsuffix F.type (trimfrom F.type "*" true)) (trimsuffix F.type (trimfrom F.type ":" true))) }}
		{{else if (contains F.type "*") ~}}
		{{> types/golang/type/new-func-field.go NAME=F.name TYP=F MOD=(ternary (trimsuffix F.type (trimfrom F.type "*" true)) (trimsuffix F.type (trimfrom F.type ":" true))) }}
		{{/if ~}}
		{{/each}}
	}
	// loop over fields looking for pointers
}


