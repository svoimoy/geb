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

	Name string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `

	Version string `json:"version" xml:"version" yaml:"version" form:"version" query:"version" `

	About string `json:"about" xml:"about" yaml:"about" form:"about" query:"about" `

	Type string `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `

	Spec map[string]interface{} `json:"spec" xml:"spec" yaml:"spec" form:"spec" query:"spec" `
}

// HOFSTADTER_BELOW
