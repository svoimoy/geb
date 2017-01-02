{{camelT NAME}}: {{> types/golang/modifiers.go MOD=MOD ~}}
{{#if IMPORT~}}
	{{IMPORT}}.
{{else ~}}
	{{> types/golang/package.go TYP ~}}
{{/if ~}}
{{#with (getsuffix (getsuffix TYP.type ":") "*") as |T|}}
{{#if (builtin T) ~}}
	{{ T ~}}
{{else ~}}
	{{camelT T ~}}
{{/if ~}}
{{/with}}
