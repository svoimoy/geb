package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
)

var (
	// A map from filename to template
	TEMPLATES map[string]*raymond.Template
)

func init() {
	TEMPLATES = make(map[string]*raymond.Template)
}

func ImportTemplateFile(filename string) error {
	return import_template("", filename)
}

func ImportTemplateFolder(folder string) error {
	// Walk the directory
	err := filepath.Walk(folder, import_template_walk_func(folder))
	if err != nil {
		return err
	}
	return nil
}

func import_template_walk_func(base_path string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		return import_template(base_path, path)
	}
}

func import_template(base_path, path string) error {
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

	// fmt.Println(" -", base_path, tpl_name)
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

	TEMPLATES[tpl_name] = tpl
	return nil
}

func add_template_helpers(tpl *raymond.Template) {

	tpl.RegisterHelper("eq", helper_eq)
	tpl.RegisterHelper("if_eq", helper_if_eq)
	tpl.RegisterHelper("if_ne", helper_if_ne)

}

func helper_eq(lhs, rhs string, options *raymond.Options) string {
	fmt.Printf("EQ: %q %q\n", lhs, rhs)
	if lhs == rhs {
		return "something"
	}
	return ""
}

func helper_if_eq(lhs, rhs string, options *raymond.Options) string {
	fmt.Printf("IF_EQ: %q %q\n", lhs, rhs)
	if lhs == rhs {
		return options.Fn()
	}
	return ""
}

func helper_if_ne(lhs, rhs string, options *raymond.Options) string {
	fmt.Printf("IF_NE: %q %q\n", lhs, rhs)
	if lhs != rhs {
		return options.Fn()
	}
	return ""
}
