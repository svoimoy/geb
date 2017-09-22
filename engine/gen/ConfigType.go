package gen

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

	Language string `json:"language" xml:"language" yaml:"language" form:"language" query:"language" `

	Dependencies Dependencies `json:"dependencies" xml:"dependencies" yaml:"dependencies" form:"dependencies" query:"dependencies" `

	TemplateConfigs []TemplateConfig `json:"template-configs" xml:"template-configs" yaml:"template-configs" form:"template-configs" query:"template-configs" `

	OutputDir string `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `
}

// HOFSTADTER_BELOW
