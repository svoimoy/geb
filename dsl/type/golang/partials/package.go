{{#if (contains TYP.type ".") ~}}
	{{#if (contains TYP.type ".views") ~}}
		{{#with (getsuffix (getsuffix TYP.type ":") "*") as |T| ~}}
			{{! check to see if package pathing still exists after trimmin views}}
			{{#if (contains (trimfrom_first T ".views" false) ".") ~}}
				{{#gettype T true ~}}
				{{trimto_last (trimfrom_last path (concat2 "." parent) false) "." false}}.{{camelT parent}}View{{! formatting ~}}
				{{/gettype ~}}
			{{else ~}}
				{{camelT (trimfrom_first T ".views" false)}}View{{! formatting ~}}
			{{/if ~}}
		{{/with ~}}
	{{else if (contains TYP.type "/") ~}}
		{{getprefix (trimto_last TYP.type "/" false) "." }}.  {{! this is just for formatting ~}}
	{{else ~}}

	{{#with (getsuffix (getsuffix TYP.type ":") "*") as |T| ~}}
		{{trimto_last (trimfrom_last T "." false) "." false ~}}.  {{! this is just for formatting ~}}
	{{/with ~}}
	{{/if~}}
{{else ~}}
{{/if~}}
