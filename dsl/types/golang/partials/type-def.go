{{#with . as |TYP|}}
{{> types/golang/package.go TYP=TYP }}
{{#with (getsuffix (getsuffix TYP.type ":") "*") as |T1|}}
{{#with (trimfrom_last T1 "." true) as |T|}}
	{{#if (builtin T) ~}}
		{{ T ~}}
	{{else ~}}
		{{camelT T ~}}
	{{/if }}
{{/with}}
{{/with}}
{{/with}}
