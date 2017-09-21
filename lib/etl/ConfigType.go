package etl

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

	InputPath string `json:"input_path" xml:"input_path" yaml:"input_path" form:"input_path" query:"input_path" `

	OutputPath string `json:"output_path" xml:"output_path" yaml:"output_path" form:"output_path" query:"output_path" `

	TemplateConfigs []gen.TemplateConfig `json:"template_configs" xml:"template_configs" yaml:"template_configs" form:"template_configs" query:"template_configs" `
}

// HOFSTADTER_BELOW
