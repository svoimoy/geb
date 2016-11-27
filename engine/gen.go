package engine

import (
	"github.com/hofstadter-io/geb/engine/project"
)

func GenerateProject(generators []string) error {

	filename := "geb.yaml"

	proj := project.NewProject()

	proj.Load(filename, generators)

	proj.Plan()

	proj.Render()

	return nil
}
