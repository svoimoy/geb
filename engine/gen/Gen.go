package gen

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"

	"github.com/hofstadter-io/geb/engine/templates"
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
func CreateFromFolder(folder string) (g *Generator, err error) {
	// HOFSTADTER_START CreateFromFolder
	g = NewGenerator()

	c, cerr := readConfigFile(filepath.Join(folder, "geb-gen.yml"))
	if cerr != nil {
		cerr = errors.Wrapf(cerr, "Error in gen.CreateFromFolder with 'geb-gen.yml' file in folder: %s\n", folder)
		c, cerr = readConfigFile(filepath.Join(folder, "geb-gen.yaml"))
		if cerr != nil {
			return nil, errors.Wrapf(cerr, "Error in gen.CreateFromFolder with 'geb-gen.yaml' file in folder: %s\n", folder)
		}
	}

	g.Config = c
	g.SourcePath = folder

	d, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "designs"))
	if err != nil {
		return nil, errors.Wrapf(err, "while reading 'designs' folder in: %s\n", folder)
	}
	g.Designs = d

	p, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "partials"))
	if err != nil {
		return nil, errors.Wrapf(err, "while reading 'partials' folder in: %s\n", folder)
	}
	g.Partials = p

	r, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "templates"))
	if err != nil {
		return nil, errors.Wrapf(err, "while reading 'repeated' folder in: %s\n", folder)
	}
	g.Templates = r

	return g, nil
	// HOFSTADTER_END   CreateFromFolder
	return
}

/*
Where's your docs doc?!
*/
func readConfigFile(filename string) (c *Config, err error) {
	// HOFSTADTER_START readConfigFile
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading generator config file: (readfile) %s\n", filename)
	}

	c = NewConfig()
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading generator config file: (unmarshal) %s\n", filename)
	}

	return c, nil
	// HOFSTADTER_END   readConfigFile
	return
}

// HOFSTADTER_BELOW
