package plan

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/templates"
	// HOFSTADTER_END   import
)

/*
Name:      plan
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
type Plan struct {
	Dsl             string                 `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `
	Gen             string                 `json:"gen" xml:"gen" yaml:"gen" form:"gen" query:"gen" `
	File            string                 `json:"file" xml:"file" yaml:"file" form:"file" query:"file" `
	Data            map[string]interface{} `json:"data" xml:"data" yaml:"data" form:"data" query:"data" `
	Template        *templates.Template    `json:"template" xml:"template" yaml:"template" form:"template" query:"template" `
	Outfile         string                 `json:"outfile" xml:"outfile" yaml:"outfile" form:"outfile" query:"outfile" `
	DslContext      interface{}            `json:"dsl-context" xml:"dsl-context" yaml:"dsl-context" form:"dsl-context" query:"dsl-context" `
	TemplateContext interface{}            `json:"template-context" xml:"template-context" yaml:"template-context" form:"template-context" query:"template-context" `
	RepeatedContext interface{}            `json:"repeated-context" xml:"repeated-context" yaml:"repeated-context" form:"repeated-context" query:"repeated-context" `
	WhenContext     interface{}            `json:"when-context" xml:"when-context" yaml:"when-context" form:"when-context" query:"when-context" `
}

func NewPlan() *Plan {
	return &Plan{
		Data: map[string]interface{}{},

		Template: templates.NewTemplate(),
	}
}

// HOFSTADTER_BELOW
