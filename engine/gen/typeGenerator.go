package gen

import (
	// HOFSTADTER_START import
	"github.ibm.com/hofstadter-io/geb/engine/templates"
	// HOFSTADTER_END   import
)

/*
Name:      generator
About:     
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
type Generator struct {
	Config *Config `json:"config" xml:"config" yaml:"config" form:"config" query:"config" `
	SourcePath string `json:"source-path" xml:"source-path" yaml:"source-path" form:"source-path" query:"source-path" `
	Templates templates.  TemplateMap `json:"templates" xml:"templates" yaml:"templates" form:"templates" query:"templates" `
	Partials templates.  TemplateMap `json:"partials" xml:"partials" yaml:"partials" form:"partials" query:"partials" `
}

func NewGenerator() *Generator {
	return &Generator{

			Config: NewConfig(),
					
			}
	// loop over fields looking for pointers
}




/*
Where's your docs doc?!
*/
func (G *Generator) MergeOverwrite(incoming *Generator)  {
	// HOFSTADTER_START MergeOverwrite
	logger.Info("Merging GEN", "existing", G.SourcePath, "incoming", incoming.SourcePath)
	for path, T := range incoming.Templates {
		_, ok := G.Templates[path]
		if ok {
			logger.Info("Overriding template", "template", path)
		} else {
			logger.Info("Adding template", "template", path)
		}
		G.Templates[path] = T
	}
	for path, P := range incoming.Partials {
		_, ok := G.Partials[path]
		if ok {
			logger.Info("Overriding partial", "partial", path)
		} else {
			logger.Info("Adding partial", "partial", path)
		}
		G.Partials[path] = P
	}
	// HOFSTADTER_END   MergeOverwrite
	return
}
/*
Where's your docs doc?!
*/
func (G *Generator) MergeSkipExisting(incoming *Generator)  {
	// HOFSTADTER_START MergeSkipExisting
	logger.Info("Merging GEN", "existing", G.SourcePath, "incoming", incoming.SourcePath)
	for path, T := range incoming.Templates {
		_, ok := G.Templates[path]
		if ok {
			logger.Info("Skipping template", "template", path)
		} else {
			logger.Info("Adding template", "template", path)
			G.Templates[path] = T
		}
	}
	for path, P := range incoming.Partials {
		_, ok := G.Partials[path]
		if ok {
			logger.Info("Skipping partial", "partial", path)
		} else {
			logger.Info("Adding partial", "partial", path)
			G.Partials[path] = P
		}
	}
	// HOFSTADTER_END   MergeSkipExisting
	return
}
/*
Where's your docs doc?!
*/
func (G *Generator) Validate() (errorReport map[string]error) {
	// HOFSTADTER_START Validate

	// HOFSTADTER_END   Validate
	return
}




// HOFSTADTER_BELOW
