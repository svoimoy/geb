package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      config
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
type Config struct {
	Name            string           `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Version         string           `json:"version" xml:"version" yaml:"version" form:"version" query:"version" `
	About           string           `json:"about" xml:"about" yaml:"about" form:"about" query:"about" `
	Type            string           `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `
	Language        string           `json:"language" xml:"language" yaml:"language" form:"language" query:"language" `
	Dependencies    Dependencies     `json:"dependencies" xml:"dependencies" yaml:"dependencies" form:"dependencies" query:"dependencies" `
	TemplateConfigs []TemplateConfig `json:"template-configs" xml:"template-configs" yaml:"template-configs" form:"template-configs" query:"template-configs" `
	OutputDir       string           `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `
}

func NewConfig() *Config {
	return &Config{
		TemplateConfigs: []TemplateConfig{},
	}
}

// HOFSTADTER_BELOW
