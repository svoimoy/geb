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

	Name string `json:"Name" xml:"Name" yaml:"Name" form:"Name" query:"Name" `

	Field string `json:"Field" xml:"Field" yaml:"Field" form:"Field" query:"Field" `

	Flatten int `json:"Flatten" xml:"Flatten" yaml:"Flatten" form:"Flatten" query:"Flatten" `

	Templates []TemplateRenderPair `json:"Templates" xml:"Templates" yaml:"Templates" form:"Templates" query:"Templates" `
}

// HOFSTADTER_BELOW
