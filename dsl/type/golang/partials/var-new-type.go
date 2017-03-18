{{#if (builtin TYP.type)}}
	// builtin
	var {{TYP.name}} {{TYP.}}

{{else}}
	// user-defined
	var input string

	/*
	{{{yaml TYP}}}
	*/

{{/if}}
