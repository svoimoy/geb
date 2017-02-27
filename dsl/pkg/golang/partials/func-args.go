{{#each . as |A|}}{{A.name}} {{A.type}}{{#if @last}}{{else}},{{/if}}{{/each}}{{! formating... ~}}
