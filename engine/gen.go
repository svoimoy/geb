package engine

import (
	"github.com/hofstadter-io/geb/engine/project"
)

func GenerateProject(filename string, generators []string) error {

	proj := project.NewProject()

	err := proj.Load(filename, generators)
	if err != nil {
		return err
	}

	err = proj.Plan()
	if err != nil {
		return err
	}

	err = proj.Render()
	if err != nil {
		return err
	}

	return nil
}
