package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      GenPair
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type GenPair struct {

	/* ORM: */

	Dsl string `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `

	Gen []string `json:"gen" xml:"gen" yaml:"gen" form:"gen" query:"gen" `

	OutputDir string `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `
}

// HOFSTADTER_BELOW
