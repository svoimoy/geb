{{#with . as |TYP|}}
{{> type/golang/modifiers.go MOD=TYP.type }}
{{> type/golang/type-def.go TYP }}
{{/with}}
