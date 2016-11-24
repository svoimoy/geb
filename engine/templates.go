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
	templates map[string]*raymond.Template
)

func init() {
	templates = make(map[string]*raymond.Template)
}

func ImportTemplateFile(filename string) error {
	return import_template(filename)
}

func ImportTemplateFolder(folder string) error {

	fmt.Println("Loading templates from: ", folder)
	// Walk the directory
	err := filepath.Walk(folder, import_template_walk_func)
	if err != nil {
		return err
	}
	return nil
}

func import_template_walk_func(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}
	if info.IsDir() {
		return nil
	}

	return import_template(path)
}

func import_template(path string) error {

	// fmt.Println(" -", path)
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

	templates[path] = tpl
	return nil
}
