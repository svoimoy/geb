`json:"{{ FIELD.name}}" xml:"{{ FIELD.name}}" yaml:"{{ FIELD.name}}" {{!formatting ~}}
form:"{{ FIELD.name}}" query:"{{ FIELD.name}}"{{!formatting ~}}
{{#if FIELD.validation }} {{> validate.go}}{{/if}} {{!formatting ~}}
{{#if FIELD.tags.orm}}{{> gorm-tags.go}}{{/if ~}} `
