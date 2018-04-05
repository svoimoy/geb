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

	Config *Config `json:"config" xml:"config" yaml:"config" form:"config" query:"config" `

	SourcePath string `json:"source-path" xml:"source-path" yaml:"source-path" form:"source-path" query:"source-path" `

	Designs templates.TemplateMap `json:"designs" xml:"designs" yaml:"designs" form:"designs" query:"designs" `

	Templates templates.TemplateMap `json:"templates" xml:"templates" yaml:"templates" form:"templates" query:"templates" `

	Partials templates.TemplateMap `json:"partials" xml:"partials" yaml:"partials" form:"partials" query:"partials" `

	NewTemplates templates.TemplateMap `json:"new-templates" xml:"new-templates" yaml:"new-templates" form:"new-templates" query:"new-templates" `
}

// HOFSTADTER_BELOW
