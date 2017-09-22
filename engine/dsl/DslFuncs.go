package dsl

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/gen"
	// HOFSTADTER_END   import
)

/*
Name:      Dsl
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewDsl() *Dsl {
	return &Dsl{

		Config: NewConfig(),

		AvailableGenerators: map[string]string{},
		Generators:          map[string]*gen.Generator{},
	}
}

/*
Where's your docs doc?!
*/
func (D *Dsl) MergeAvailable(incoming *Dsl) {
	// HOFSTADTER_START MergeAvailable
	logger.Info("Merging Available", "existing", D.Config.Name, "incoming", incoming.Config.Name)
	for path, G := range incoming.Generators {
		_, ok := D.Generators[path]
		if !ok {
			logger.Info("Adding Generator", "generator", path)
			D.Generators[path] = G
		}
	}
	// HOFSTADTER_END   MergeAvailable
	return
}

/*
Where's your docs doc?!
*/
func (D *Dsl) MergeOverwrite(incoming *Dsl) {
	// HOFSTADTER_START MergeOverwrite
	logger.Info("Merging DSLs", "existing", D.Config.Name, "incoming", incoming.Config.Name)
	for path, G := range incoming.Generators {
		existing, ok := D.Generators[path]
		if ok {
			logger.Info("Merging Gen")
			existing.MergeOverwrite(G)
			D.Generators[path] = existing
		} else {
			logger.Info("Adding Gen")
			D.Generators[path] = G
		}
	}
	// HOFSTADTER_END   MergeOverwrite
	return
}

/*
Where's your docs doc?!
*/
func (D *Dsl) MergeSkipExisting(incoming *Dsl) {
	// HOFSTADTER_START MergeSkipExisting
	logger.Info("Merging DSLs", "existing", D.Config.Name, "incoming", incoming.Config.Name)
	for path, G := range incoming.Generators {
		existing, ok := D.Generators[path]
		if ok {
			logger.Info("Merging Gen")
			existing.MergeSkipExisting(G)
			D.Generators[path] = existing
		} else {
			logger.Info("Adding Gen")
			D.Generators[path] = G
		}
	}
	// HOFSTADTER_END   MergeSkipExisting
	return
}

/*
Where's your docs doc?!
*/
func (D *Dsl) Validate() (errorReport map[string]error) {
	// HOFSTADTER_START Validate

	// HOFSTADTER_END   Validate
	return
}

// HOFSTADTER_BELOW
