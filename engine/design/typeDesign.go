package design

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v1"

	"github.ibm.com/hofstadter-io/dotpath"
	// HOFSTADTER_END   import
)

/*
Name:      design
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
type Design struct {
	Custom map[string]interface{} `json:"custom" xml:"custom" yaml:"custom" form:"custom" query:"custom" `
	Dsl    map[string]interface{} `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `
	Pkg    map[string]interface{} `json:"pkg" xml:"pkg" yaml:"pkg" form:"pkg" query:"pkg" `
	Proj   map[string]interface{} `json:"proj" xml:"proj" yaml:"proj" form:"proj" query:"proj" `
	Type   map[string]interface{} `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `
}

func NewDesign() *Design {
	return &Design{
		Custom: map[string]interface{}{},
		Dsl:    map[string]interface{}{},
		Pkg:    map[string]interface{}{},
		Proj:   map[string]interface{}{},
		Type:   map[string]interface{}{},
	}
	// loop over fields looking for pointers
}

/*
Where's your docs doc?!
*/
func (D *Design) ImportDesignFile(filename string) (err error) {
	// HOFSTADTER_START ImportDesignFile
	logger.Info("Importing Design filename:", "filename", filename)
	return D.importDesign(filepath.Dir(filename), filename)
	// HOFSTADTER_END   ImportDesignFile
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) ImportDesignFolder(folder string) (err error) {
	// HOFSTADTER_START ImportDesignFolder
	logger.Info("Importing Design folder: " + folder)

	// Make sure the folder exists
	_, err = os.Lstat(folder)
	if err != nil {
		return errors.Wrapf(err, "error lstat'n path in utils.ResolvePath\n")
	}

	// local walk function closure
	import_design_walk_func := func(path string, info os.FileInfo, err error) error {
		local_d := D
		if err != nil {
			return nil
		}
		if info.IsDir() || !(strings.Contains(info.Name(), ".yml") || strings.Contains(info.Name(), ".yaml")) {
			return nil
		}

		return local_d.importDesign(folder, path)
	}

	// Walk the directory
	err = filepath.Walk(folder, import_design_walk_func)
	if err != nil {
		return errors.Wrap(err, "in design.CreateFromFolder: "+folder+"\n")
	}
	return nil
	// HOFSTADTER_END   ImportDesignFolder
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) Get(path string) (object interface{}, err error) {
	// HOFSTADTER_START Get
	return D.GetByPath(path)
	// HOFSTADTER_END   Get
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) GetByPath(path string) (object interface{}, err error) {
	// HOFSTADTER_START GetByPath
	paths := strings.Split(path, ".")
	if len(paths) < 1 {
		return nil, errors.New("Bad path supplied: " + path)
	}

	if len(paths) == 1 {
		switch paths[0] {
		case "custom":
			return D.Custom, nil

		case "dsl":
			return D.Dsl, nil

		case "pkg":
			return D.Pkg, nil

		case "proj":
			return D.Proj, nil

		case "type":
			return D.Type, nil

		default:
			return nil, errors.New("Unknown path start for design: " + paths[0])

		}
	}
	P, rest := paths[0], paths[1:]

	switch P {
	case "custom":
		return dotpath.GetByPathSlice(rest, D.Custom, true)

	case "dsl":
		return dotpath.GetByPathSlice(rest, D.Dsl, true)

	case "pkg":
		return dotpath.GetByPathSlice(rest, D.Pkg, true)

	case "proj":
		return dotpath.GetByPathSlice(rest, D.Proj, true)

	case "type":
		return dotpath.GetByPathSlice(rest, D.Type, true)

	default:
		return nil, errors.New("Unknown path start for design: " + P)

	}
	// HOFSTADTER_END   GetByPath
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) Validate() (errorReport map[string]error) {
	// HOFSTADTER_START Validate

	// HOFSTADTER_END   Validate
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) importDesign(basePath string, designPath string) (err error) {
	// HOFSTADTER_START importDesign
	logger.Info("  - file: " + designPath)
	// fmt.Println(" -", designPath)
	top_level := make(map[string]interface{})
	raw_data, err := ioutil.ReadFile(designPath)
	if err != nil {
		return errors.Wrap(err, "in design.import_design (read file): "+designPath+"\n")
	}
	err = yaml.Unmarshal([]byte(raw_data), &top_level)
	if err != nil {
		return errors.Wrap(err, "in design.import_design (yaml unmarshal): "+designPath+"\n")
	}

	rel_file, err := filepath.Rel(basePath, designPath)
	if err != nil {
		return errors.Wrap(err, "in design.import_design (rel filepath): "+designPath+"\n")
	}

	rel_path := filepath.Dir(rel_file)
	if rel_path[0] == '.' {
		rel_path = rel_path[1:]
	}

	// get list of all top level DSL entries
	for dsl, val := range top_level {
		data := val.(map[interface{}]interface{})
		err = D.storeDesign(rel_path, dsl, data)
		if err != nil {
			return errors.Wrap(err, "in design.import_design (store design): "+designPath+"\n")
		}
	}

	return nil
	// HOFSTADTER_END   importDesign
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) storeDesign(relativePath string, dsl string, design interface{}) (err error) {
	// HOFSTADTER_START storeDesign
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
			logger.Debug("dotpath for name", "dname", dname, "err", err, "elem", elem)

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
			D.storeTypeDesign(relativePath, ename, elem)

		}

	case "type":
		D.storeTypeDesign(relativePath, name, design)

	case "pkg":
		D.storePackageDesign(relativePath, name, design)

	case "custom":
		D.storeCustomDesign(relativePath, name, design)

	default:
		D.storeDslDesign(relativePath, dsl, name, design)
	}
	return nil
	// HOFSTADTER_END   storeDesign
	return
}

// HOFSTADTER_BELOW
