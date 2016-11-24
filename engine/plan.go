package engine

import "fmt"

type Plan struct {
	template string
	design   interface{}
	outfile  string
}

func MakeRenderingPlans(generators []string) ([]Plan, error) {
	plans := make([]Plan, 0)

	// loop over types
	for key, _ := range DESIGN.types {
		fmt.Println("Making plan for design:", key)
	}

	// loop over designs
	for key, _ := range DESIGN.dsls {
		fmt.Println("Making plan for design:", key)
	}

	// loop over custom
	for key, _ := range DESIGN.custom {
		fmt.Println("Making plan for custom:", key)
	}

	return plans, nil
}
