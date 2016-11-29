package design

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v1"
)

type DesignData map[interface{}]interface{}
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
	return d.import_design(filename)
}

func (d *Design) ImportDesignFolder(folder string) error {

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
	switch dsl {
	case "api", "cli":
		_, ok := design["name"]
		if !ok {
			return errors.New("field 'name' missing from " + dsl + " dsl")
		}
		d.Dsl[dsl] = design
		logger.Info("  found dsl", "dsl", dsl, "name", design["name"])

	case "type":
		iname, ok := design["name"]
		if !ok {
			return errors.New("field 'name' missing from TYPE dsl")
		}
		name, ok := iname.(string)
		if !ok {
			return errors.New("field 'name' is not a string in TYPE " + fmt.Sprint(iname))
		}
		d.Type[name] = design
		logger.Info("  found type", "type", dsl, "name", design["name"])

	default:
		d.Custom[dsl] = design
		logger.Info("  found custom", "name", dsl)
	}

	return nil
}
