package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v1"
)

type Design struct {
	custom map[string]DesignData
	dsls   map[string]DesignData
	types  map[string]DesignData
}

type DesignData map[interface{}]interface{}

var (
	DESIGN Design
)

func init() {
	DESIGN.custom = make(map[string]DesignData)
	DESIGN.dsls = make(map[string]DesignData)
	DESIGN.types = make(map[string]DesignData)
}

func ImportDesignFile(filename string) error {
	return import_design(filename)
}

func ImportDesignFolder(folder string) error {

	fmt.Println("Loading designs from: ", folder)
	// Walk the directory
	err := filepath.Walk(folder, import_design_walk_func)
	if err != nil {
		return err
	}
	return nil
}

func import_design_walk_func(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}
	if info.IsDir() || !(strings.Contains(info.Name(), ".yml") || strings.Contains(info.Name(), ".yaml")) {
		return nil
	}

	return import_design(path)
}

func import_design(path string) error {
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
		err = store_design(dsl, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func store_design(dsl string, design DesignData) error {
	switch dsl {
	case "api", "cli":
		_, ok := design["name"]
		if !ok {
			return errors.New("field 'name' missing from " + dsl + " dsl")
		}
		DESIGN.dsls[dsl] = design
		// fmt.Println("    ", dsl, design["name"])
	case "type":
		iname, ok := design["name"]
		if !ok {
			return errors.New("field 'name' missing from TYPE dsl")
		}
		name, ok := iname.(string)
		if !ok {
			return errors.New("field 'name' is not a string in TYPE " + fmt.Sprint(iname))
		}
		DESIGN.types[name] = design
		// fmt.Println("    ", dsl, design["name"])

	default:
		DESIGN.custom[dsl] = design
		// fmt.Println("    ", dsl, "data")
	}

	return nil
}
