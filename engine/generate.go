package engine

import (
	"errors"
	"io/ioutil"
)

func Generate(generators []string, outdir string) error {

	// make rendering plans
	plans, err := MakeRenderingPlans(generators)
	if err != nil {
		return err
	}

	// render designs
	for _, plan := range plans {
		err = RenderPlan(plan)
		if err != nil {
			return err
		}
	}

	return nil
}

func RenderPlan(plan Plan) error {
	result, err := RenderTemplate(plan.template, plan.design)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(plan.outfile, []byte(result), 0644)
	if err != nil {
		return err
	}

	return nil
}

func RenderTemplate(path string, design interface{}) (string, error) {
	tpl, ok := templates[path]
	if !ok {
		return "", errors.New("Unknown template: " + path)
	}
	return tpl.Exec(design)
}
