package templates

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      template
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
type Template struct {
	ID   string `json:"id" xml:"id" yaml:"id" form:"id" query:"id" `
	Name string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Data string `json:"data" xml:"data" yaml:"data" form:"data" query:"data" `
}
