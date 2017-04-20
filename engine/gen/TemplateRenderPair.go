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
	When    string   `json:"when" xml:"when" yaml:"when" form:"when" query:"when" `
	Unless  string   `json:"unless" xml:"unless" yaml:"unless" form:"unless" query:"unless" `
	Field   string   `json:"field" xml:"field" yaml:"field" form:"field" query:"field" `
	Flatten int      `json:"flatten" xml:"flatten" yaml:"flatten" form:"flatten" query:"flatten" `
	In      string   `json:"in" xml:"in" yaml:"in" form:"in" query:"in" `
	Out     string   `json:"out" xml:"out" yaml:"out" form:"out" query:"out" `
	Extra   []string `json:"extra" xml:"extra" yaml:"extra" form:"extra" query:"extra" `
}

func NewTemplateRenderPair() *TemplateRenderPair {
	return &TemplateRenderPair{
		Extra: []string{},
	}
}

// HOFSTADTER_BELOW
