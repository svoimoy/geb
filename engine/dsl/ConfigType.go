package dsl

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Config
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Config struct {

	/* ORM: */

	Name string `json:"Name" xml:"Name" yaml:"Name" form:"Name" query:"Name" `

	Version string `json:"Version" xml:"Version" yaml:"Version" form:"Version" query:"Version" `

	About string `json:"About" xml:"About" yaml:"About" form:"About" query:"About" `

	Type string `json:"Type" xml:"Type" yaml:"Type" form:"Type" query:"Type" `

	Spec map[string]interface{} `json:"Spec" xml:"Spec" yaml:"Spec" form:"Spec" query:"Spec" `
}

// HOFSTADTER_BELOW
