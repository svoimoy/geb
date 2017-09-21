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

	Files map[string]interface{} `json:"Files" xml:"Files" yaml:"Files" form:"Files" query:"Files" `

	Proj map[string]interface{} `json:"Proj" xml:"Proj" yaml:"Proj" form:"Proj" query:"Proj" `

	Data map[string]interface{} `json:"Data" xml:"Data" yaml:"Data" form:"Data" query:"Data" `

	Type map[string]interface{} `json:"Type" xml:"Type" yaml:"Type" form:"Type" query:"Type" `

	Pkg map[string]interface{} `json:"Pkg" xml:"Pkg" yaml:"Pkg" form:"Pkg" query:"Pkg" `

	Dsl map[string]interface{} `json:"Dsl" xml:"Dsl" yaml:"Dsl" form:"Dsl" query:"Dsl" `

	Custom map[string]interface{} `json:"Custom" xml:"Custom" yaml:"Custom" form:"Custom" query:"Custom" `

	Extra map[string]interface{} `json:"Extra" xml:"Extra" yaml:"Extra" form:"Extra" query:"Extra" `
}

// HOFSTADTER_BELOW
