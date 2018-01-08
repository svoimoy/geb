package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      TemplateConfig
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type TemplateConfig struct {

	/* ORM: */

	Name string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `

	Field string `json:"field" xml:"field" yaml:"field" form:"field" query:"field" `

	Flatten int `json:"flatten" xml:"flatten" yaml:"flatten" form:"flatten" query:"flatten" `

	Templates []TemplateRenderPair `json:"templates" xml:"templates" yaml:"templates" form:"templates" query:"templates" `

	StaticFiles []StaticFilesConfig `json:"static-files" xml:"static-files" yaml:"static-files" form:"static-files" query:"static-files" `
}

// HOFSTADTER_BELOW
