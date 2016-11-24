package engine

import "errors"

func Generate(generators, outdir string) error {

	// make plan

	// render designs

	return nil
}

func RenderTemplate(path string, design interface{}) (string, error) {
	tpl, ok := templates[path]
	if !ok {
		return "", errors.New("Unknown template: " + path)
	}
	return tpl.Exec(design)
}
