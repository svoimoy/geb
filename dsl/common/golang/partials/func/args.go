{{#each . as |A|}}{{camel A.name}} {{> type/golang/type.go TYP=A}}{{#if @last}}{{else}},{{/if}}{{/each}}{{! formating... ~}}
