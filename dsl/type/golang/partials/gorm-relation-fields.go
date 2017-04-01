{{#if TYP.relations}}
// Relations ----
// Should be using some type lookup and package resolution type things
{{#each TYP.relations as |REL|}}

{{#if (eq REL.relation "has-many")}}
	// has-many
	{{camelT REL.name}} []{{camelT REL.type}} {{> type/golang/tags.go FIELD=REL}}
{{else if (eq REL.relation "belongs-to")}}
	// belongs-to
	{{camelT REL.name}} {{camelT REL.type}} {{> type/golang/tags.go FIELD=REL}}
	{{camelT REL.name}}ID string // the type should be inferred from {{camelT REL.type}}.{{camelT REL.foreign-key}}
{{else if (eq REL.relation "many-to-many")}}
	// many-to-many
	{{camelT REL.name}} []{{camelT REL.type}} {{> type/golang/tags.go FIELD=REL}}
{{else}}
	// unknown REL.relation: '{{REL.relation}}'
{{/if}}

{{/each}}
{{/if}}
