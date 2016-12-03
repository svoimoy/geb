{{#with type.count}}
package main

// Name:    {{name}}
// Version: {{version}}
type {{title name}} struct {
{{#each fields}}
{{#if private}}
private: '{{name}}'
{{else}}
public:  '{{name}}' {{> tags.go}}
{{/if}}
{{/each}}
}

{{/with}}
