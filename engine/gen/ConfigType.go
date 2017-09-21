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

	Name string `json:"Name" xml:"Name" yaml:"Name" form:"Name" query:"Name" `

	Version string `json:"Version" xml:"Version" yaml:"Version" form:"Version" query:"Version" `

	About string `json:"About" xml:"About" yaml:"About" form:"About" query:"About" `

	Type string `json:"Type" xml:"Type" yaml:"Type" form:"Type" query:"Type" `

	Language string `json:"Language" xml:"Language" yaml:"Language" form:"Language" query:"Language" `

	Dependencies Dependencies `json:"Dependencies" xml:"Dependencies" yaml:"Dependencies" form:"Dependencies" query:"Dependencies" `

	TemplateConfigs []TemplateConfig `json:"TemplateConfigs" xml:"TemplateConfigs" yaml:"TemplateConfigs" form:"TemplateConfigs" query:"TemplateConfigs" `

	OutputDir string `json:"OutputDir" xml:"OutputDir" yaml:"OutputDir" form:"OutputDir" query:"OutputDir" `
}

// HOFSTADTER_BELOW
