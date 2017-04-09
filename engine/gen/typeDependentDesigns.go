package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      dependent-designs
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
type DependentDesigns struct {
	Name  string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Field string `json:"field" xml:"field" yaml:"field" form:"field" query:"field" `
	When  string `json:"when" xml:"when" yaml:"when" form:"when" query:"when" `
	In    string `json:"in" xml:"in" yaml:"in" form:"in" query:"in" `
	Out   string `json:"out" xml:"out" yaml:"out" form:"out" query:"out" `
}

func NewDependentDesigns() *DependentDesigns {
	return &DependentDesigns{}
}

// HOFSTADTER_BELOW
