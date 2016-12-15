{{#with . as |F|}}
{{#if F.private}}
{{camel F.name}} {{>type.go F.type ~}}
{{else}}
{{camelT F.name}} {{>type.go F.type}} {{> tags.go F ~}}
{{/if}}
{{/with}}

