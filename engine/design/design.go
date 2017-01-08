package design

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.ibm.com/hofstadter-io/dotpath"
	"gopkg.in/yaml.v1"
	// HOFSTADTER_END   import
)

// Name:      design
// Namespace: engine.design
// Version:   0.0.1

type Design struct {
	Proj   map[string]interface{} `json:"proj" xml:"proj" yaml:"proj" form:"proj" query:"proj" `
	Pkg    map[string]interface{} `json:"pkg" xml:"pkg" yaml:"pkg" form:"pkg" query:"pkg" `
	Type   map[string]interface{} `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `
	Dsl    map[string]interface{} `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `
	Custom map[string]interface{} `json:"custom" xml:"custom" yaml:"custom" form:"custom" query:"custom" `
}

func NewDesign() *Design {
	return &Design{
		Proj:   map[string]interface{}{},
		Pkg:    map[string]interface{}{},
		Type:   map[string]interface{}{},
		Dsl:    map[string]interface{}{},
		Custom: map[string]interface{}{},
	}
	// loop over fields looking for pointers
}

// HOFSTADTER_BELOW

func New() *Design {
	return NewDesign()
}

func CreateFromFolder(folder string) (*Design, error) {
	d := NewDesign()
	err := d.ImportDesignFolder(folder)
	if err != nil {
		return nil, errors.Wrap(err, "in design.CreateFromFolder: "+folder+"\n")
	}
	return d, nil
}

func (d *Design) ImportDesignFile(filename string) error {
	logger.Info("Importing Design fileame:", "filename", filename)
	return d.import_design(filepath.Dir(filename), filename)
}

func (d *Design) ImportDesignFolder(folder string) error {
	logger.Info("Importing Design folder: " + folder)

	// Make sure the folder exists
	_, err := os.Lstat(folder)
	if err != nil {
		return errors.Wrapf(err, "error lstat'n path in utils.ResolvePath\n")
	}

	// local walk function closure
	import_design_walk_func := func(path string, info os.FileInfo, err error) error {
		local_d := d
		if err != nil {
			return nil
		}
		if info.IsDir() || !(strings.Contains(info.Name(), ".yml") || strings.Contains(info.Name(), ".yaml")) {
			return nil
		}

		return local_d.import_design(folder, path)
	}

	// Walk the directory
	err = filepath.Walk(folder, import_design_walk_func)
	if err != nil {
		return errors.Wrap(err, "in design.CreateFromFolder")
	}
	return nil
}

func (d *Design) import_design(base_folder, path string) error {
	logger.Info("  - file: " + path)
	// fmt.Println(" -", path)
	top_level := make(map[string]interface{})
	raw_data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "in design.import_design")
	}
	err = yaml.Unmarshal([]byte(raw_data), &top_level)
	if err != nil {
		return errors.Wrap(err, "in design.import_design")
	}

	rel_file, err := filepath.Rel(base_folder, path)
	if err != nil {
		return errors.Wrap(err, "in design.import_design")
	}

	rel_path := filepath.Dir(rel_file)
	if rel_path[0] == '.' {
		rel_path = rel_path[1:]
	}

	// get list of all top level DSL entries
	for dsl, val := range top_level {
		data := val.(map[interface{}]interface{})
		err = d.store_design(rel_path, dsl, data)
		if err != nil {
			return errors.Wrap(err, "in design.import_design")
		}
	}

	return nil
}

func (d *Design) store_design(relative_path, dsl string, design interface{}) error {
	logger.Info("store_design: " + dsl)

	dname, err := dotpath.Get("name", design, true)
	logger.Debug("dotpath for name", "dname", dname, "err", err, "design", design)

	// Everything must have a name!
	name := ""
	switch D := design.(type) {

	case map[string]interface{}:
		iname, ok := D["name"]
		if !ok {
			return errors.New("Top-level definition '" + dsl + "' missing required field 'name'.")
		}
		tmp, ok := iname.(string)
		if !ok {
			return errors.New("Top-level definition '" + dsl + "' field 'name' is not a string.")
		}
		name = tmp

	case map[interface{}]interface{}:
		iname, ok := D["name"]
		if !ok {
			return errors.New("Top-level definition '" + dsl + "' missing required field 'name'.")
		}
		tmp, ok := iname.(string)
		if !ok {
			return errors.New("Top-level definition '" + dsl + "' field 'name' is not a string.")
		}
		name = tmp

	default:
		return errors.New("Top-level definition '" + dsl + "' must be a map type.\nTry adding a single top-level entry with the rest under it.")

	}

	if name == "" {
		return errors.New("Top-level definition '" + dsl + "' field 'name' is empty.")
	}

	switch dsl {
	case "type-list":
		var t_list []interface{}
		switch D := design.(type) {
		case map[string]interface{}:
			tmp_list, ok := D["list"].([]interface{})
			if !ok {
				return errors.New("Top-level type-list does not have a 'list' or is not an array of objects in '" + " design: " + fmt.Sprint(design))
			}
			t_list = tmp_list

		case map[interface{}]interface{}:
			tmp_list, ok := D["list"].([]interface{})
			if !ok {
				return errors.New("Top-level type-list does not have a 'list' or is not an array of objects in '" + " design: " + fmt.Sprint(design))
			}
			t_list = tmp_list

		default:
			return errors.New("Type-list definition '" + dsl + "' must be a map type.\nTry adding a single top-level entry with the rest under it.")

		}
		for _, elem := range t_list {
			ename := ""
			// check that we have a name, and possibly overwrite namespace
			dname, err := dotpath.Get("name", elem, true)
			logger.Warn("dotpath for name", "dname", dname, "err", err, "elem", elem)

			switch E := elem.(type) {

			case map[string]interface{}:
				iname, ok := E["name"]
				if !ok {
					return errors.New("Type-list definition '" + name + "' missing required field 'name'.")
				}
				tmp, ok := iname.(string)
				if !ok {
					return errors.New("Type-list definition '" + name + "' field 'name' is not a string.")
				}
				ename = tmp

			case map[interface{}]interface{}:
				iname, ok := E["name"]
				if !ok {
					return errors.New("Type-list definition '" + name + "' missing required field 'name'.")
				}
				tmp, ok := iname.(string)
				if !ok {
					return errors.New("Type-list definition '" + name + "' field 'name' is not a string.")
				}
				ename = tmp

			default:
				return errors.New("Type-list definition '" + dsl + "' is not a map[string]")

			}
			d.store_type_design(relative_path, ename, elem)

		}

	case "type":
		d.store_type_design(relative_path, name, design)

	case "pkg":
		d.store_pkg_design(relative_path, name, design)

	case "custom":
		d.store_custom_design(relative_path, name, design)

	default:
		d.store_dsl_design(relative_path, dsl, name, design)
	}
	return nil
}
