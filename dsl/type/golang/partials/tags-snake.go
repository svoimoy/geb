`json:"{{snake FIELD.name}}" xml:"{{snake FIELD.name}}" yaml:"{{snake FIELD.name}}" {{!formatting ~}}
form:"{{snake FIELD.name}}" query:"{{snake FIELD.name}}"{{!formatting ~}}
{{#if FIELD.validation }} {{> validate.go}}{{/if}} {{!formatting ~}}
{{#if FIELD.tags.orm}}{{> gorm-tags.go}}{{/if ~}} `
