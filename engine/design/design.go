package design

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
	"gopkg.in/yaml.v1"
	// HOFSTADTER_END   import
)

// Name:      design
// Namespace: engine.design
// Version:   0.0.1

type Design struct {
	Proj   map[string]interface{} `json:"proj" xml:"proj" yaml:"proj" form:"proj" query:"proj" `
	Type   map[string]interface{} `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `
	Dsl    map[string]interface{} `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `
	Custom map[string]interface{} `json:"custom" xml:"custom" yaml:"custom" form:"custom" query:"custom" `
}

func NewDesign() *Design {
	return &Design{
		Proj:   map[string]interface{}{},
		Type:   map[string]interface{}{},
		Dsl:    map[string]interface{}{},
		Custom: map[string]interface{}{},
	}
	// loop over fields looking for pointers
}

/*
fields:
- name: proj
  type: map:interface{}
- name: type
  type: map:interface{}
- name: dsl
  type: map:interface{}
- name: custom
  type: map:interface{}
name: design
namespace: engine.design
version: 0.0.1

*/

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
	return d.import_design(filename)
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

		return local_d.import_design(path)
	}

	// Walk the directory
	err = filepath.Walk(folder, import_design_walk_func)
	if err != nil {
		return errors.Wrap(err, "in design.CreateFromFolder")
	}
	return nil
}

func (d *Design) import_design(path string) error {
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

	// get list of all top level DSL entries
	for dsl, val := range top_level {
		data := val.(map[interface{}]interface{})
		err = d.store_design(dsl, data)
		if err != nil {
			return errors.Wrap(err, "in design.import_design")
		}
	}

	return nil
}

func (d *Design) store_design(dsl string, design interface{}) error {
	logger.Info("    - storing: " + dsl)
	logger.Debug("        data:", "design", design, "dsl", dsl)

	// Everything must have a name!
	name := ""
	NS := "" // namespace from type list
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
		ns, ok := D["namespace"].(string)
		if ok {
			NS = ns
		}

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
		ns, ok := D["namespace"].(string)
		if ok {
			NS = ns
		}

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
				return errors.New("Top-level type-list is not a list or has non-map objects '" + " design: " + fmt.Sprint(design))
			}
			t_list = tmp_list

		case map[interface{}]interface{}:
			tmp_list, ok := D["list"].([]interface{})
			if !ok {
				return errors.New("Top-level type-list is not a list or has non-map objects '" + " design: " + fmt.Sprint(design))
			}
			t_list = tmp_list

		default:
			return errors.New("Type-list definition '" + dsl + "' must be a map type.\nTry adding a single top-level entry with the rest under it.")

		}
		for _, elem := range t_list {
			ename := ""
			switch E := elem.(type) {

			case map[string]interface{}:
				if NS != "" {
					E["namespace"] = NS
				}
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
				if NS != "" {
					E["namespace"] = NS
				}
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

			logger.Info("      - storing type", "type", dsl, "ename", ename, "elem", elem, "NS", NS)
			fields := strings.Split(NS, ".")
			F0 := fields[0]
			logger.Info("Fields", "fields", fields)

			if L := len(fields); L > 0 {

				insert := make(map[string]interface{})
				dd_map := insert
				for i, F := range fields {
					tmp := make(map[string]interface{})
					logger.Info("FIELD_INDEX A", "i", i, "F", F, "map", dd_map, "tmp", tmp, "insert", insert)
					dd_map[F] = tmp
					logger.Info("FIELD_INDEX B", "i", i, "F", F, "map", dd_map, "tmp", tmp, "insert", insert)
					dd_map = tmp
					logger.Info("FIELD_INDEX C", "i", i, "F", F, "map", dd_map, "tmp", tmp, "insert", insert)
				}

				dd_map[ename] = elem
				logger.Info("Design", "ename", ename, "design", design, "map", dd_map, "insert", insert)

				if curr, ok := d.Type[F0]; !ok {
					d.Type[F0] = insert[F0]
					logger.Debug("d.TypeList new L>1", "data", d.Type)
				} else {
					logger.Info("...", "curr", curr, "d.Type", d.Type, "update", insert)
					merged, merr := utils.Merge(d.Type, insert)
					if merr != nil {
						return errors.Wrap(merr, "in store_design typelist.loop")
					}
					logger.Info("result...", "merged", merged)
					d.Type[F0] = merged.(map[string]interface{})[F0]
					logger.Debug("d.TypeList merge L>1", "data", d.Type)
				}
				logger.Debug("       - " + F0)

			} else {
				if curr, ok := d.Type[name]; !ok {
					d.Type[name] = elem
					logger.Debug("d.TypeList new L<2", "data", d.Type)
				} else {
					merged, merr := utils.Merge(curr, elem)
					if merr != nil {
						return errors.Wrap(merr, "in store_design typelist")
					}
					d.Type[name] = merged
					logger.Debug("d.TypeList merge L<2", "data", d.Type)
				}
			}
		}

	case "type":

		logger.Info("    - storing type", "type", dsl, "name", name, "design", design, "NS", NS)
		fields := strings.Split(NS, ".")
		F0 := fields[0]
		// FL := fields[len(fields)-1]
		logger.Info("Fields", "fields", fields)

		if L := len(fields); L > 0 {

			insert := make(map[string]interface{})
			dd_map := insert
			for i, F := range fields {
				tmp := make(map[string]interface{})
				logger.Info("FIELD_INDEX A", "i", i, "F", F, "map", dd_map, "tmp", tmp, "insert", insert)
				dd_map[F] = tmp
				logger.Info("FIELD_INDEX B", "i", i, "F", F, "map", dd_map, "tmp", tmp, "insert", insert)
				dd_map = tmp
				logger.Info("FIELD_INDEX C", "i", i, "F", F, "map", dd_map, "tmp", tmp, "insert", insert)
			}

			dd_map[name] = design
			logger.Info("Design", "name", name, "design", design, "map", dd_map, "insert", insert)

			if curr, ok := d.Type[F0]; !ok {
				d.Type[F0] = insert[F0]
				logger.Debug("d.Type new L>1", "data", d.Type)
			} else {
				logger.Info("merge...", "curr", curr, "d.Type", d.Type, "update", insert)
				merged, merr := utils.Merge(d.Type, insert)
				if merr != nil {
					return errors.Wrap(merr, "in store_design type.loop")
				}
				logger.Info("result...", "merged", merged)
				d.Type[F0] = merged.(map[string]interface{})[F0]
				logger.Debug("d.Type merge L>1", "data", d.Type)
			}
			logger.Debug("       - " + F0)

		} else {
			if curr, ok := d.Type[name]; !ok {
				d.Type[name] = design
				logger.Debug("d.Type new L<2", "data", d.Type)
			} else {
				merged, merr := utils.Merge(curr, design)
				if merr != nil {
					return errors.Wrap(merr, "in store_design type")
				}
				d.Type[name] = merged
				logger.Debug("d.Type merge L<2", "data", d.Type)
			}
		}

	case "custom":
		if curr, ok := d.Type[name]; !ok {
			d.Custom[name] = design
		} else {
			merged, merr := utils.Merge(curr, design)
			if merr != nil {
				return errors.Wrap(merr, "in store_design custom")
			}
			d.Custom[name] = merged
		}
		logger.Debug("    - storing custom", "type", dsl, "name", name)

	default:
		logger.Debug("    - storing dsl", "type", dsl, "name", name)
		fields := strings.Split(dsl, "/")
		logger.Debug("Fields", "fields", fields)

		if L := len(fields); L > 0 {

			dd_map := design
			logger.Debug("       - "+name, "L", L)

			for i := L - 1; i > 0; i-- {
				curr := fields[i]
				logger.Debug("       - "+curr, "L", L, "i", i)
				tmp_map := make(map[string]interface{})
				tmp_map[curr] = dd_map
				dd_map = tmp_map
			}

			if curr, ok := d.Dsl[fields[0]]; !ok {
				d.Dsl[fields[0]] = dd_map
			} else {
				merged, merr := utils.Merge(curr, dd_map)
				if merr != nil {
					return errors.Wrap(merr, "in store_design dsl.loop (default)")
				}
				d.Dsl[fields[0]] = merged
			}
			logger.Debug("       - " + fields[0])

		} else {
			if curr, ok := d.Dsl[dsl]; !ok {
				d.Dsl[dsl] = design
			} else {
				merged, merr := utils.Merge(curr, design)
				if merr != nil {
					return errors.Wrap(merr, "in store_design dsl (default)")
				}
				d.Dsl[dsl] = merged
			}
		}

	}

	return nil
}
