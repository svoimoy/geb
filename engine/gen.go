package engine

import (
	"fmt"
	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/project"
)

func GenerateProject(filename string, generators []string) error {

	proj := project.NewProject()

	fmt.Println("Loading...")
	err := proj.Load(filename, generators)
	if err != nil {
		return errors.Wrapf(err, "While generating project: %s %v\n", filename, generators)
	}

	fmt.Println("Planning...")
	err = proj.Plan()
	if err != nil {
		return errors.Wrapf(err, "While planing project: %s %v\n", filename, generators)
	}

	fmt.Println("Rendering...")
	err = proj.Render()
	if err != nil {
		return errors.Wrapf(err, "While rendering project: %s %v\n", filename, generators)
	}

	return nil
}
