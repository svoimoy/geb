package dsl

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hofstadter-io/geb/engine/gen"
	"gopkg.in/yaml.v1"
)

type Dsl struct {
	Name    string
	Version string
	About   string

	SourcePath          string
	AvailableGenerators []string

	Generators map[string]gen.Generator
}

func NewDsl() *Dsl {
	return &Dsl{
		AvailableGenerators: []string{},
		Generators:          map[string]gen.Generator{},
	}
}

func LoadDsl(folder string) (*Dsl, error) {
	D, err := ReadDslFile(filepath.Join(folder, "dsl.yaml"))
	if err != nil {
		return nil, err
	}

	D.SourcePath = folder
	return D, nil
}

func ReadDslFile(filename string) (*Dsl, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	d := NewDsl()
	err = yaml.Unmarshal(data, d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func FindAvailable(folder string) (map[string]*Dsl, error) {
	logger.Info("Searching for DSLs", "folder", folder)
	dsls := map[string]*Dsl{}
	var curr_dsl *Dsl

	// local walk function closure
	walk_func := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() || !(strings.Contains(info.Name(), ".yml") || strings.Contains(info.Name(), ".yaml")) {
			return nil
		}

		dir, fn := filepath.Split(path)

		if fn == "geb-dsl.yml" || fn == "geb-dsl.yaml" {
			rel, err := filepath.Rel(folder, dir)
			if err != nil {
				return err
			}
			curr_dsl = NewDsl()
			curr_dsl.Name = rel
			curr_dsl.SourcePath = dir
			dsls[rel] = curr_dsl
			logger.Info("  found DSL", "name", rel)
		}

		if fn == "geb-gen.yml" || fn == "geb-gen.yaml" {
			rel, err := filepath.Rel(curr_dsl.SourcePath, dir)
			if err != nil {
				return err
			}
			logger.Info("    found gen: ", "dsl", curr_dsl.Name, "name", rel)
			curr_dsl.AvailableGenerators = append(curr_dsl.AvailableGenerators, rel)
		}

		return nil
	}

	// Walk the directory
	err := filepath.Walk(folder, walk_func)
	if err != nil {
		return nil, err
	}

	return dsls, nil
}
