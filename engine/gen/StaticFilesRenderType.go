package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      StaticFilesRender
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type StaticFilesRender struct {

	/* ORM: */

	When string `json:"when" xml:"when" yaml:"when" form:"when" query:"when" `

	Unless string `json:"unless" xml:"unless" yaml:"unless" form:"unless" query:"unless" `

	Field string `json:"field" xml:"field" yaml:"field" form:"field" query:"field" `

	Files []string `json:"files" xml:"files" yaml:"files" form:"files" query:"files" `

	Ignores []string `json:"ignores" xml:"ignores" yaml:"ignores" form:"ignores" query:"ignores" `
}

// HOFSTADTER_BELOW
