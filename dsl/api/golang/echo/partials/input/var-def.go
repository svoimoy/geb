{{#with . as |CURR|}}
	// CURR: {{CURR}}
	// RC:   {{RC}}
	{{#get_elem_by_name input "" data=@root.type}}
	{{#if (contains input ".views.")}}
	// This is the view:  '{{input}}'
	{{#fields}}
		{{#if local-ref}}
		{{#get_elem_by_name (join3 "." RC.name "fields" local-ref) "" data=@root.type}}
		// local-ref:  {{../local-ref}} -> {{{name}}}
		//-- var input_{{name}} {{type}}
		
		
		{{/get_elem_by_name}}
		{{else if remote-ref}}
		// remote-ref:  {{remote-ref}}
		{{else}}
		// Error: unknown view field type
		{{/if}}
	{{/fields}}
	{{else}}
	// This is the input:  '{{input}}'
	//-- var input_{{name}} {{type}}
	{{/if}}
	{{/get_elem_by_name}}
{{/with}}
