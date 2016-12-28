{{#with . as |T|}}
{{#if (contains T "array:") ~}} []{{/if ~}}
{{#if (contains T "map:") ~}} map[string]{{/if ~}}
{{#if (contains T "*") ~}}*{{/if ~}}
{{#with (getsuffix (getsuffix T ":") "*") as |TYP|}}
{{#if (contains T ".") ~}}
	{{#each (rsublist (split TYP ".") 0 2 ) ~}}
    {{#if @last ~}}
	{{#if (builtin .) ~}}
		{{ . ~}}
	{{else ~}}
		{{camelT . ~}}
	{{/if }}
	  {{else ~}}{{.}}.{{/if ~}}
	{{/each}}
{{else}}
	{{#if (builtin TYP) ~}}
		{{ TYP ~}}
	{{else ~}}
		{{camelT TYP ~}}
	{{/if }}
{{/if ~}}
{{/with}}
{{/with}}
