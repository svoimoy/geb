var {{NAME}} {{> type/golang/modifiers.go MOD=MOD ~}}
{{#if IMPORT~}}
{{IMPORT}}.
{{else ~}}
{{> type/golang/package.go TYP ~}}
{{/if~}}
{{camelT TYP.parent}}View{{camelT TYP.name}}
