{{#with . as |TYP|}}
{{#with (getsuffix (getsuffix TYP.type ":") "*") as |T1|}}
{{#with (trimto_last T1 "." false) as |T|}}
	{{#if (builtin T) ~}}
		{{ T ~}}
	{{else ~}}
		{{camelT T ~}}
	{{/if }}
{{/with}}
{{/with}}
{{/with}}
