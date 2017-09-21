package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Generator
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewGenerator() *Generator {
	return &Generator{

		Config: NewConfig(),
	}
}

/*
Where's your docs doc?!
*/
func (G *Generator) MergeOverwrite(incoming *Generator) {
	// HOFSTADTER_START MergeOverwrite
	logger.Info("Merging GEN", "existing", G.SourcePath, "incoming", incoming.SourcePath)

	for path, T := range incoming.Designs {
		_, ok := G.Designs[path]
		if ok {
			logger.Info("Overriding design", "design", path)
		} else {
			logger.Info("Adding design", "design", path)
		}
		G.Designs[path] = T
	}

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
func (G *Generator) MergeSkipExisting(incoming *Generator) {
	// HOFSTADTER_START MergeSkipExisting
	logger.Info("Merging GEN", "existing", G.SourcePath, "incoming", incoming.SourcePath)

	for path, T := range incoming.Designs {
		_, ok := G.Designs[path]
		if ok {
			logger.Info("Skipping design", "design", path)
		} else {
			logger.Info("Adding design", "design", path)
			G.Designs[path] = T
		}
	}

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
