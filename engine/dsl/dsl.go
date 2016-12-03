package dsl

import (
	"github.com/pkg/errors"
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
	Type    string

	SourcePath          string
	AvailableGenerators map[string]string

	Generators map[string]*gen.Generator
}

func NewDsl() *Dsl {
	return &Dsl{
		AvailableGenerators: map[string]string{},
		Generators:          map[string]*gen.Generator{},
	}
}

func CreateFromFolder(folder string) (*Dsl, error) {
	D, err := ReadDslFile(filepath.Join(folder, "geb-dsl.yml"))
	if err != nil {
		err = errors.Wrapf(err, "Error in dsl.CreateFromFolder with 'geb-dsl.yml' file in folder: %s\n", folder)
		D, err = ReadDslFile(filepath.Join(folder, "geb-dsl.yaml"))
		if err != nil {
			return nil, errors.Wrapf(err, "Error in dsl.CreateFromFolder with 'geb-dsl.yaml' file in folder: %s\n", folder)
		}
	}

	D.SourcePath = folder
	return D, nil
}

func ReadDslFile(filename string) (*Dsl, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading dsl config file: (readfile) %s\n", filename)
	}

	d := NewDsl()
	err = yaml.Unmarshal(data, d)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading dsl config file: (unmarshal) %s\n", filename)
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

		logger.Debug("Walking:  " + path)

		if info.IsDir() {
			// check to see if there is a geb file (due to walking in alpha order)
			// if so, just rewrite the path and things should work out
			// we will make sure we aren't overwriting the second time witht the DSL
			fns := []string{
				"geb-dsl.yml",
				"geb-dsl.yaml",
				"geb-gen.yml",
				"geb-gen.yaml",
			}

			geb_fn := ""
			for _, fn := range fns {
				g_fn := filepath.Join(path, fn)
				_, g_err := os.Lstat(g_fn)
				if g_err == nil {
					geb_fn = g_fn
					break
				}
			}

			if geb_fn != "" {
				path = geb_fn
			} else {
				// skip this directory
				return nil
			}
		} else if !(strings.Contains(info.Name(), ".yml") || strings.Contains(info.Name(), ".yaml")) {
			// only interested in yaml files (actually geb files) for making decisions
			return nil
		}

		dir, fn := filepath.Split(path)

		if fn == "geb-dsl.yml" || fn == "geb-dsl.yaml" {
			rel, err := filepath.Rel(folder, dir)
			if err != nil {
				return errors.Wrapf(err, "In dsl.FindAvailable:  %s %s\n", folder, dir)
			}
			if _, ok := dsls[rel]; ok {
				// already discovered this dsl
				return nil
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
				return errors.Wrapf(err, "In dsl.FindAvailable:  %s %s\n", curr_dsl.SourcePath, dir)
			}
			if _, ok := curr_dsl.AvailableGenerators[rel]; ok {
				// already discovered this dsl
				return nil
			}
			logger.Info("    generator: ", "dsl", curr_dsl.Name, "name", rel)
			curr_dsl.AvailableGenerators[rel] = rel
		}

		return nil
	}

	info, err := os.Lstat(folder)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		dir, err := os.Readlink(folder)
		if err != nil {
			return nil, err
		}
		folder = dir
	}
	// Walk the directory
	err = filepath.Walk(folder, walk_func)
	if err != nil {
		return nil, err
	}

	return dsls, nil
}

func (D *Dsl) MergeAvailable(fresh *Dsl) {
	logger.Info("Merging Available", "existing", D.Name, "fresh", fresh.Name)
	for path, G := range fresh.Generators {
		_, ok := D.Generators[path]
		if !ok {
			logger.Info("Adding Generator", "generator", path)
			D.Generators[path] = G
		}
	}
}

func (D *Dsl) MergeSkipExisting(fresh *Dsl) {
	logger.Info("Merging DSLs", "existing", D.Name, "fresh", fresh.Name)
	for path, G := range fresh.Generators {
		existing, ok := D.Generators[path]
		if ok {
			logger.Info("Merging Gen")
			G.MergeSkipExisting(existing)
			D.Generators[path] = G
		} else {
			logger.Info("Adding Gen")
			D.Generators[path] = G
		}
	}
}

func (D *Dsl) MergeOverwrite(fresh *Dsl) {
	logger.Info("Merging DSLs", "existing", D.Name, "fresh", fresh.Name)
	for path, G := range fresh.Generators {
		existing, ok := D.Generators[path]
		if ok {
			logger.Info("Merging Gen")
			existing.MergeOverwrite(G)
			D.Generators[path] = existing
		} else {
			logger.Info("Adding Gen")
			D.Generators[path] = G
		}
	}
}
