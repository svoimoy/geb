package dsl

// package

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func CreateFromFolder(folder string) (d *Dsl, err error) {
	// HOFSTADTER_START CreateFromFolder
	d = NewDsl()

	C, err := ReadConfigFile(filepath.Join(folder, "geb-dsl.yml"))
	if err != nil {
		// logger.Info("error reading, geb-dsl.yml, trying geb-dsl.yaml", "err", err)
		err = errors.Wrapf(err, "Error in dsl.CreateFromFolder with 'geb-dsl.yml' file in folder: %s\n", folder)
		C2, err2 := ReadConfigFile(filepath.Join(folder, "geb-dsl.yaml"))
		if err2 != nil {
			err2 = errors.Wrap(err, "error reading geb-dsl.yaml, giving up.\n")
			return nil, errors.Wrapf(err2, "Error in dsl.CreateFromFolder with 'geb-dsl.yaml' file in folder: %s\n", folder)
		}
		err = nil
		C = C2
	}
	d.Config = C

	d.SourcePath = folder
	return d, nil
	// HOFSTADTER_END   CreateFromFolder
	return
}

/*
Where's your docs doc?!
*/
func FindAvailable(folder string) (ds map[string]*Dsl, err error) {
	// HOFSTADTER_START FindAvailable
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
				"geb-gen.yaml"}

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
			curr_dsl.Config.Name = rel
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
			logger.Info("    generator: ", "dsl", curr_dsl.Config.Name, "name", rel)
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
	// HOFSTADTER_END   FindAvailable
	return
}

// HOFSTADTER_BELOW
