package gen

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/templates"
	// HOFSTADTER_END   import
)

/*
Name:      Generator
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Generator struct {

	/* ORM: */

	Config *Config `json:"Config" xml:"Config" yaml:"Config" form:"Config" query:"Config" `

	SourcePath string `json:"SourcePath" xml:"SourcePath" yaml:"SourcePath" form:"SourcePath" query:"SourcePath" `

	Designs templates.TemplateMap `json:"Designs" xml:"Designs" yaml:"Designs" form:"Designs" query:"Designs" `

	Templates templates.TemplateMap `json:"Templates" xml:"Templates" yaml:"Templates" form:"Templates" query:"Templates" `

	Partials templates.TemplateMap `json:"Partials" xml:"Partials" yaml:"Partials" form:"Partials" query:"Partials" `
}

// HOFSTADTER_BELOW
