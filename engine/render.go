package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func RenderPlans(plans []Plan, outdir string) error {

	// render designs
	for _, plan := range plans {
		err := RenderPlan(plan, outdir)
		if err != nil {
			return err
		}
	}

	return nil
}

func RenderPlan(plan Plan, outdir string) error {
	if plan.template != "" {
		fmt.Printf("Rendering single template: %q", plan.template)
		result, err := RenderTemplate(plan.template, plan.design)
		if err != nil {
			return err
		}

		fmt.Println("  - redering to ", plan.outfile, len(result))
		err = WriteResults(plan.outfile, result)
		if err != nil {
			return err
		}
	}

	if len(plan.templates) > 0 {
		for _, path := range plan.templates {
			result, err := RenderTemplate(path, plan.design)
			if err != nil {
				return err
			}

			outfile := filepath.Join(outdir, path)
			fmt.Println("  - redering to ", outfile, len(result))
			err = WriteResults(outfile, result)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func RenderTemplate(path string, design interface{}) (string, error) {
	tpl, ok := TEMPLATES[path]
	if !ok {
		return "", errors.New("Unknown template: " + path)
	}
	return tpl.Exec(design)
}

func WriteResults(filename, content string) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
