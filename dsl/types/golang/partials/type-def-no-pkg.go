{{#with . as |T|}}
/* T: {{{pprint T}}} */
{{#with (trimto_last (getsuffix (getsuffix T ":") "*") "." false) as |TYP|}}
/* TYP: {{{pprint TYP}}} */
	{{#if (builtin TYP) ~}}
		{{ TYP ~}}
	{{else ~}}
		{{camelT TYP ~}}
	{{/if }}
{{/with}}
{{/with}}
