{{#if typename}}
	// gettype: {{{typename}}}
	{{#gettype typename true as |TYP|}}
		{{> type/golang/var-new-type.go TYP=TYP }}
	{{/gettype}}

{{else}}
	{{> type/golang/var-new-type.go TYP=TYP }}
{{/if}}

