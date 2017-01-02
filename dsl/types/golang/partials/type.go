{{#with . as |TYP|}}
{{> types/golang/modifiers.go MOD=TYP.type }}
{{> types/golang/type-def.go TYP }}
{{/with}}
