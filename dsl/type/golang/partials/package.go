{{#if (contains type ".") ~}}
	{{#if (hassuffix type ".views") ~}}
		{{#gettype (trimsuffix type ".views") true ~}}
		{{snake (getsuffix (trimto type (concat2 "." parent) false) ".") ~}}.  {{! this is just for formatting ~}}
		{{/gettype ~}}
	{{else if (contains type "/") ~}}
		{{getprefix (trimto_last type "/" false) "." }}.  {{! this is just for formatting ~}}
	{{else ~}}
	{{#with (getsuffix (getsuffix type ":") "*") as |T| ~}}
		{{getsuffix (trimto_last T "." true) "." ~}}.  {{! this is just for formatting ~}}
	{{/with ~}}
	{{/if~}}
{{else ~}}
{{/if~}}
