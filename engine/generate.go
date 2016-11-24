package engine

import "errors"

func Generate(generators, outdir string) error {

	return nil
}

func RenderTemplate(path string) (string, error) {
	tpl, ok := templates[path]
	if !ok {
		return "", errors.New("Unknown template: " + path)
	}
	return tpl.Exec(DESIGN)
}
