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

	When string `json:"when" xml:"when" yaml:"when" form:"when" query:"when" `

	Unless string `json:"unless" xml:"unless" yaml:"unless" form:"unless" query:"unless" `

	Field string `json:"field" xml:"field" yaml:"field" form:"field" query:"field" `

	Flatten int `json:"flatten" xml:"flatten" yaml:"flatten" form:"flatten" query:"flatten" `

	In string `json:"in" xml:"in" yaml:"in" form:"in" query:"in" `

	Out string `json:"out" xml:"out" yaml:"out" form:"out" query:"out" `

	Extra []string `json:"extra" xml:"extra" yaml:"extra" form:"extra" query:"extra" `
}

// HOFSTADTER_BELOW
