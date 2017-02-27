{{#if (contains path ".") ~}}
	{{#if (hassuffix path ".views") ~}}
		{{#gettype (trimsuffix path ".views") true ~}}
		{{snake (getsuffix (trimto path (concat2 "." parent) false) ".") ~}}.  {{! this is just for formatting ~}}
		{{/gettype ~}}
	{{else if (contains type "/") ~}}
		{{getprefix (trimfrom_last type "/" true) "." }}.  {{! this is just for formatting ~}}
	{{else ~}}
	{{#with (getsuffix (getsuffix type ":") "*") as |T| ~}}
		{{getsuffix (trimto_last T "." false) "." ~}}.  {{! this is just for formatting ~}}
	{{/with ~}}
	{{/if~}}
{{else ~}}
{{/if~}}
