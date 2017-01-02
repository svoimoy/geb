{{#if TYP.package ~}}
	{{snake (TYP.package)}}.
{{else if (contains TYP.type ".") ~}}
	{{#if (hassuffix path ".views") ~}}
		{{#gettype (trimsuffix path ".views") true ~}}
			{{#if namespace}}{{snake namespace}}.{{/if ~}}
		{{/gettype ~}}
	{{else ~}}
		{{#gettype path true ~}}
			{{#if namespace}}{{snake namespace}}.{{/if ~}}
		{{/gettype ~}}
	{{/if~}}
{{/if~}}

