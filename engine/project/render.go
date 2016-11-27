package project

import (
	"github.com/aymerick/raymond"
	"github.com/hofstadter-io/geb/engine/gen"
)

func (P *Project) Render() error {
	logger.Warn("Project.Render() TBD")
	return nil
}

func RenderTemplate(template *gen.Template, design interface{}) (string, error) {
	tpl := (*raymond.Template)(template)
	return tpl.Exec(design)
}
