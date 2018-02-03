package plan

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/templates"
	// HOFSTADTER_END   import
)

/*
Name:      Plan
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Plan struct {

	/* ORM: */

	Dsl string `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `

	Gen string `json:"gen" xml:"gen" yaml:"gen" form:"gen" query:"gen" `

	File string `json:"file" xml:"file" yaml:"file" form:"file" query:"file" `

	StaticContent []byte `json:"static-content" xml:"static-content" yaml:"static-content" form:"static-content" query:"static-content" `

	Data map[string]interface{} `json:"data" xml:"data" yaml:"data" form:"data" query:"data" `

	Template *templates.Template `json:"template" xml:"template" yaml:"template" form:"template" query:"template" `

	Outfile string `json:"outfile" xml:"outfile" yaml:"outfile" form:"outfile" query:"outfile" `

	DslContext interface{} `json:"dsl-context" xml:"dsl-context" yaml:"dsl-context" form:"dsl-context" query:"dsl-context" `

	TemplateContext interface{} `json:"template-context" xml:"template-context" yaml:"template-context" form:"template-context" query:"template-context" `

	RepeatedContext interface{} `json:"repeated-context" xml:"repeated-context" yaml:"repeated-context" form:"repeated-context" query:"repeated-context" `

	WhenContext interface{} `json:"when-context" xml:"when-context" yaml:"when-context" form:"when-context" query:"when-context" `
}

// HOFSTADTER_BELOW
