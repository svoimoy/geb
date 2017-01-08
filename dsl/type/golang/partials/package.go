{{#if TYP.package ~}}
	{{snake (TYP.package)}}.
{{else if (contains TYP.type ".") ~}}
	{{#if (hassuffix path ".views") ~}}
		{{#gettype (trimsuffix path ".views") true ~}}
			{{#if namespace}}{{snake namespace}}.{{/if ~}}
		{{/gettype ~}}
	{{else if (contains TYP.type "/") ~}}
		{{getprefix (trimfrom_last TYP.type "/" true) "." }}.  {{! this is just for formatting ~}}
	{{else if path ~}}
		{{#gettype path true ~}}
			{{#if namespace}}{{snake namespace}}.{{/if ~}}
		{{/gettype ~}}
	{{ else ~}}
	{{#with (getsuffix (getsuffix TYP.type ":") "*") as |T| ~}}
		{{getsuffix (trimto_last T "." false) "." ~}}.  {{! this is just for formatting ~}}
	{{/with ~}}
	{{/if~}}
{{ else ~}}
{{/if~}}
