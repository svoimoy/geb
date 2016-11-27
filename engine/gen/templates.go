package gen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
)

type Template raymond.Template
type TemplateMap map[string]*Template

func NewTemplateMap() TemplateMap {
	return make(map[string]*Template)
}

func CreateTemplateMapFromFolder(folder string) (TemplateMap, error) {
	M := NewTemplateMap()
	err := M.ImportFromFolder(folder)
	if err != nil {
		return nil, err
	}
	return M, nil
}

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
		return err
	}

	add_template_helpers(tpl)

	M[tpl_name] = (*Template)(tpl)
	return nil
}

func add_template_helpers(tpl *raymond.Template) {

	tpl.RegisterHelper("eq", helper_eq)
	tpl.RegisterHelper("if_eq", helper_if_eq)
	tpl.RegisterHelper("if_ne", helper_if_ne)

}

func helper_eq(lhs, rhs string, options *raymond.Options) string {
	str := fmt.Sprintf("EQ: %q %q\n", lhs, rhs)
	logger.Debug(str)
	if lhs == rhs {
		return "something"
	}
	return ""
}

func helper_if_eq(lhs, rhs string, options *raymond.Options) string {
	str := fmt.Sprintf("IF_EQ: %q %q\n", lhs, rhs)
	logger.Debug(str)
	if lhs == rhs {
		return options.Fn()
	}
	return ""
}

func helper_if_ne(lhs, rhs string, options *raymond.Options) string {
	str := fmt.Sprintf("IF_NE: %q %q\n", lhs, rhs)
	logger.Debug(str)
	if lhs != rhs {
		return options.Fn()
	}
	return ""
}
