package engine

import (
	"errors"
	"fmt"

	"github.com/aymerick/raymond"
	"github.com/ryanuber/go-glob"
)

type Plan struct {
	template  string
	design    interface{}
	outfile   string
	templates []string
}

/*
Force the issue of folder structure, dsl designs, plan making,
and extensiblity based on nameing conventions

Make sveral examples and move to actual apps.

Start with "count-service" and "hofstadter.io"
*/

func MakeRenderingPlans(generators []string, design interface{}, templates map[string]*raymond.Template) ([]Plan, error) {
	plans := make([]Plan, 0)

	fmt.Println("Planning generators:")
	for _, g := range generators {
		var plan Plan
		plan.design = design
		spath := g + "*"

		for path, _ := range templates {
			found := glob.Glob(spath, path)
			//	fmt.Println(spath, path, found)
			if found {
				// fmt.Println("   -", path)
				plan.templates = append(plan.templates, path)
			}
		}

		if len(plan.templates) == 0 {
			return nil, errors.New("No templates found for generator: " + g)
		}
		fmt.Printf("  %4d files:  %s\n", len(plan.templates), g)
		plans = append(plans, plan)
	}

	/*
		// loop over types
		for key, _ := range design.types {
			fmt.Println("Making plan for design:", key)
		}

		// loop over designs
		for key, _ := range design.dsls {
			fmt.Println("Making plan for design:", key)
		}

		// loop over custom
		for key, _ := range design.custom {
			fmt.Println("Making plan for custom:", key)
		}
	*/

	return plans, nil
}
