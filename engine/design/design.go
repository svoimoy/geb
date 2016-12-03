package design

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v1"
)

type DesignData map[interface{}]interface{}

// these two should be in with the next map at load time
// DesignRoot string // Folder from design list in Project
// 	SourceFile string // The actual design the file
type DesignDataMap map[string]DesignData

type Design struct {
	Proj   map[string]interface{}
	Type   DesignDataMap
	Dsl    DesignDataMap
	Custom DesignDataMap

	RepeatedContext interface{}
}

func NewDesign() *Design {
	return &Design{
		Proj:   make(map[string]interface{}),
		Type:   make(DesignDataMap),
		Dsl:    make(DesignDataMap),
		Custom: make(DesignDataMap),
	}
}
func CreateFromFolder(folder string) (*Design, error) {
	d := NewDesign()
	d.ImportDesignFolder(folder)
	return d, nil
}

func (d *Design) ImportDesignFile(filename string) error {
	logger.Info("Importing Design fileame:", "filename", filename)
	return d.import_design(filename)
}

func (d *Design) ImportDesignFolder(folder string) error {
	logger.Info("Importing Design folder: " + folder)

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
	err := filepath.Walk(folder, import_design_walk_func)
	if err != nil {
		return err
	}
	return nil
}

func (d *Design) import_design(path string) error {
	logger.Info("  - file: " + path)
	// fmt.Println(" -", path)
	top_level := make(map[string]interface{})
	raw_data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(raw_data), &top_level)
	if err != nil {
		return err
	}

	// get list of all top level DSL entries
	for dsl, val := range top_level {
		data := val.(map[interface{}]interface{})
		err = d.store_design(dsl, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Design) store_design(dsl string, design DesignData) error {
	logger.Info("    - storing: " + dsl)
	logger.Debug("        data:", "design", design, "dsl", dsl)

	// Everything must have a name!
	iname, ok := design["name"]
	if !ok {
		return errors.New("Top-level definition '" + dsl + "' missing required field 'name'")
	}
	name, ok := iname.(string)
	if !ok {
		return errors.New("Top-level definition '" + dsl + "' field 'name' is not a string")
	}

	switch dsl {
	case "type":
		_, overwrite := d.Type[name]
		d.Type[name] = design
		logger.Debug("    - storing type", "type", dsl, "name", name, "overwrite", overwrite)

	case "custom":
		_, overwrite := d.Custom[name]
		d.Custom[name] = design
		logger.Debug("    - storing custom", "type", dsl, "name", name, "overwrite", overwrite)

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
				tmp_map := make(DesignData)
				tmp_map[curr] = dd_map
				dd_map = tmp_map
			}

			d.Dsl[fields[0]] = dd_map
			logger.Debug("       - " + fields[0])

		} else {
			d.Dsl[dsl] = design
		}

	}

	return nil
}
