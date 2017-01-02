{{#with . as |T|}}
{{> types/golang/modifiers.go MOD=T }}
{{> types/golang/type-def.go T }}
{{/with}}
