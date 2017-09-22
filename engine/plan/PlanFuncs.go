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

func NewPlan() *Plan {
	return &Plan{
		Data: map[string]interface{}{},

		Template: templates.NewTemplate(),
	}
}

// HOFSTADTER_BELOW
