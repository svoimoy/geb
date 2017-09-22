package design

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Design
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Design struct {

	/* ORM: */

	Files map[string]interface{} `json:"files" xml:"files" yaml:"files" form:"files" query:"files" `

	Proj map[string]interface{} `json:"proj" xml:"proj" yaml:"proj" form:"proj" query:"proj" `

	Data map[string]interface{} `json:"data" xml:"data" yaml:"data" form:"data" query:"data" `

	Type map[string]interface{} `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `

	Pkg map[string]interface{} `json:"pkg" xml:"pkg" yaml:"pkg" form:"pkg" query:"pkg" `

	Dsl map[string]interface{} `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `

	Custom map[string]interface{} `json:"custom" xml:"custom" yaml:"custom" form:"custom" query:"custom" `

	Extra map[string]interface{} `json:"extra" xml:"extra" yaml:"extra" form:"extra" query:"extra" `
}

// HOFSTADTER_BELOW
