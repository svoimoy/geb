package etl

import (
	// HOFSTADTER_START import
	"github.ibm.com/hofstadter-io/geb/engine/gen"
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
	InputPath       string               `json:"input_path" xml:"input_path" yaml:"input_path" form:"input_path" query:"input_path" `
	OutputPath      string               `json:"output_path" xml:"output_path" yaml:"output_path" form:"output_path" query:"output_path" `
	TemplateConfigs []gen.TemplateConfig `json:"template_configs" xml:"template_configs" yaml:"template_configs" form:"template_configs" query:"template_configs" `
}

func NewConfig() *Config {
	return &Config{
		TemplateConfigs: []gen.TemplateConfig{},
	}
}

// HOFSTADTER_BELOW
