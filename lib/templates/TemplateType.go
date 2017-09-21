package templates

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Template
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Template struct {

	/* ORM: */

	Id string `json:"id" xml:"id" yaml:"id" form:"id" query:"id" `

	Name string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `

	Data string `json:"data" xml:"data" yaml:"data" form:"data" query:"data" `
}

// HOFSTADTER_BELOW
