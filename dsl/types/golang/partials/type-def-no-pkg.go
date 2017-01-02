{{#with . as |T|}}
{{#with (getsuffix (getsuffix T ":") "*") as |T1|}}
{{#with (trimfrom_last T1 "." true) as |TYP|}}
	{{#if (builtin TYP) ~}}
		{{ TYP ~}}
	{{else ~}}
		{{camelT TYP ~}}
	{{/if }}
{{/with}}
{{/with}}
{{/with}}
