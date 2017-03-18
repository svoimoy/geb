{{#if (contains MOD "channel") ~}} chan {{/if ~}}
{{#if (contains MOD "array") ~}} []{{/if ~}}
{{#if (contains MOD "map") ~}} map[string]{{/if ~}}
{{#if (contains MOD "*") ~}}*{{/if ~}}
