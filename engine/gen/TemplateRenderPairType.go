package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      TemplateRenderPair
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type TemplateRenderPair struct {

	/* ORM: */

	When string `json:"When" xml:"When" yaml:"When" form:"When" query:"When" `

	Unless string `json:"Unless" xml:"Unless" yaml:"Unless" form:"Unless" query:"Unless" `

	Field string `json:"Field" xml:"Field" yaml:"Field" form:"Field" query:"Field" `

	Flatten int `json:"Flatten" xml:"Flatten" yaml:"Flatten" form:"Flatten" query:"Flatten" `

	In string `json:"In" xml:"In" yaml:"In" form:"In" query:"In" `

	Out string `json:"Out" xml:"Out" yaml:"Out" form:"Out" query:"Out" `

	Extra []string `json:"Extra" xml:"Extra" yaml:"Extra" form:"Extra" query:"Extra" `
}

// HOFSTADTER_BELOW
