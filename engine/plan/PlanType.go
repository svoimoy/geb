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

	Dsl string `json:"Dsl" xml:"Dsl" yaml:"Dsl" form:"Dsl" query:"Dsl" `

	Gen string `json:"Gen" xml:"Gen" yaml:"Gen" form:"Gen" query:"Gen" `

	File string `json:"File" xml:"File" yaml:"File" form:"File" query:"File" `

	Data map[string]interface{} `json:"Data" xml:"Data" yaml:"Data" form:"Data" query:"Data" `

	Template *templates.Template `json:"Template" xml:"Template" yaml:"Template" form:"Template" query:"Template" `

	Outfile string `json:"Outfile" xml:"Outfile" yaml:"Outfile" form:"Outfile" query:"Outfile" `

	DslContext interface{} `json:"DslContext" xml:"DslContext" yaml:"DslContext" form:"DslContext" query:"DslContext" `

	TemplateContext interface{} `json:"TemplateContext" xml:"TemplateContext" yaml:"TemplateContext" form:"TemplateContext" query:"TemplateContext" `

	RepeatedContext interface{} `json:"RepeatedContext" xml:"RepeatedContext" yaml:"RepeatedContext" form:"RepeatedContext" query:"RepeatedContext" `

	WhenContext interface{} `json:"WhenContext" xml:"WhenContext" yaml:"WhenContext" form:"WhenContext" query:"WhenContext" `
}

// HOFSTADTER_BELOW
