package templates

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
type Template struct {
	*raymond.Template
}

func NewTemplate() *Template {
	return &Template{}
}

func NewMap() TemplateMap {
	return NewTemplateMap()
}

func (template *Template) Render(design interface{}) (string, error) {
	tpl := (*raymond.Template)(template.Template)
	return tpl.Exec(design)
}

type TemplateMap map[string]*Template

func NewTemplateMap() TemplateMap {
	return make(map[string]*Template)
}

// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func RenderTemplate(template *Template, design interface{}) (output string, err error) {
	// HOFSTADTER_START RenderTemplate
	tpl := (*raymond.Template)(template.Template)
	return tpl.Exec(design)
	// HOFSTADTER_END   RenderTemplate
	return
}

/*
Where's your docs doc?!
*/
func AddHelpersToRaymond(tpl *Template) {
	// HOFSTADTER_START AddHelpersToRaymond
	rtpl := (*raymond.Template)(tpl.Template)
	addTemplateHelpers(rtpl)
	// HOFSTADTER_END   AddHelpersToRaymond
	return
}

/*
Where's your docs doc?!
*/
func AddHelpersToTemplate(tpl *Template) {
	// HOFSTADTER_START AddHelpersToTemplate
	rtpl := (*raymond.Template)(tpl.Template)
	addTemplateHelpers(rtpl)
	// HOFSTADTER_END   AddHelpersToTemplate
	return
}

/*
Where's your docs doc?!
*/
func CreateTemplateFromFile(filename string) (tpl *Template, err error) {
	// HOFSTADTER_START CreateTemplateFromFile

	// HOFSTADTER_END   CreateTemplateFromFile
	return
}

/*
Where's your docs doc?!
*/
func CreateTemplateMapFromFolder(folder string) (tplMap TemplateMap, err error) {
	// HOFSTADTER_START CreateTemplateMapFromFolder
	tplMap = NewTemplateMap()
	err = tplMap.ImportFromFolder(folder)
	if err != nil {
		return nil, errors.Wrapf(err, "while importing %s\n", folder)
	}
	return tplMap, nil
	// HOFSTADTER_END   CreateTemplateMapFromFolder
	return
}

// HOFSTADTER_BELOW

func (M TemplateMap) ImportTemplateFile(filename string) error {
	return M.import_template("", filename)
}

func (M TemplateMap) ImportFromFolder(folder string) error {
	import_template_walk_func := func(base_path string) filepath.WalkFunc {
		return func(path string, info os.FileInfo, err error) error {
			local_m := M
			if err != nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}

			return local_m.import_template(base_path, path)
		}
	}

	// Walk the directory
	err := filepath.Walk(folder, import_template_walk_func(folder))
	if err != nil {
		return err
	}
	return nil
}

func (M TemplateMap) import_template(base_path, path string) error {
	tpl_name := path
	L := len(base_path)
	if L > 0 {
		// should handle trailing slashes better here
		if base_path[L-1] == '/' {
			tpl_name = path[L:]
		} else {
			tpl_name = path[L+1:]
		}
	}

	logger.Info("    found template", "path", base_path, "name", tpl_name)
	raw_template, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	source := string(raw_template)

	// parse template
	tpl, err := raymond.Parse(source)
	if err != nil {
		return errors.Wrapf(err, "While parsing file: %s\n", tpl_name)
	}

	addTemplateHelpers(tpl)

	M[tpl_name] = &Template{tpl}
	return nil
}
