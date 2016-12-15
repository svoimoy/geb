{{#with . as |T|}}
{{#if (contains T "array:")}}
	{{#if (builtin (getsuffix T ":"))}}
	[]{{getsuffix T ":" ~}}
	{{else}}
		[]{{camelT (getsuffix T ":") ~}}
	{{/if}}
{{else if (contains T "map:")}}
	{{#if (builtin T)}}
		map[string]{{getsuffix T ":" ~}}
	{{else}}
		map[string]{{camelT (trimprefix T ":") ~}}
	{{/if}}
{{else if (builtin T)}}
	{{ T ~}}
{{else}}
	{{camelT T ~}}
{{/if}}
{{/with}}
