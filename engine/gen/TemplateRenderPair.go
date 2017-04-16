package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      template-render-pair
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
type TemplateRenderPair struct {
	When   string `json:"when" xml:"when" yaml:"when" form:"when" query:"when" `
	Unless string `json:"unless" xml:"unless" yaml:"unless" form:"unless" query:"unless" `
	In     string `json:"in" xml:"in" yaml:"in" form:"in" query:"in" `
	Out    string `json:"out" xml:"out" yaml:"out" form:"out" query:"out" `
}

func NewTemplateRenderPair() *TemplateRenderPair {
	return &TemplateRenderPair{}
}

// HOFSTADTER_BELOW
