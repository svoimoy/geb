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

/*
Where's your docs doc?!
*/
type TemplateRenderPair struct {
	In string `json:"in" xml:"in" yaml:"in" form:"in" query:"in" `
	Out string `json:"out" xml:"out" yaml:"out" form:"out" query:"out" `
}

func NewTemplateRenderPair() *TemplateRenderPair {
	return &TemplateRenderPair{
	}
	// loop over fields looking for pointers
}








// HOFSTADTER_BELOW
