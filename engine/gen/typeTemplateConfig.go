package gen

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      template-config
About:     
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
type TemplateConfig struct {
	Name string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Field string `json:"field" xml:"field" yaml:"field" form:"field" query:"field" `
	Templates []TemplateRenderPair `json:"templates" xml:"templates" yaml:"templates" form:"templates" query:"templates" `
}

func NewTemplateConfig() *TemplateConfig {
	return &TemplateConfig{
Templates: []TemplateRenderPair{},
			}
	// loop over fields looking for pointers
}








// HOFSTADTER_BELOW
