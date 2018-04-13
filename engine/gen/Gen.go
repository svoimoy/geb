package gen

import (
	// HOFSTADTER_START import

	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"

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
	logger.Info("!!!!!!!!!! GOT HERE GENERATOR")

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
	logger.Info("!!!!!!!!!!FOUND in generator", "g", g, "config", *g.Config)

	// Read all geb-gen*.yaml files
	matches, err := filepath.Glob(filepath.Join(folder, "geb-gen-*.yaml"))
	if err != nil {
		return nil, errors.Wrapf(err, "while globing for 'geb-gen-*.yaml' in folder: %s\n", folder)
	}
	matches2, err := filepath.Glob(filepath.Join(folder, "geb-gen/*.yaml"))
	if err != nil {
		return nil, errors.Wrapf(err, "while globing for 'geb-gen-*.yaml' in folder: %s\n", folder)
	}
	matches = append(matches, matches2...)
	matches3, err := filepath.Glob(filepath.Join(folder, "geb-gen/*/*.yaml"))
	if err != nil {
		return nil, errors.Wrapf(err, "while globing for 'geb-gen-*.yaml' in folder: %s\n", folder)
	}
	matches = append(matches, matches3...)
	matches4, err := filepath.Glob(filepath.Join(folder, "geb-gen/*/*/*.yaml"))
	if err != nil {
		return nil, errors.Wrapf(err, "while globing for 'geb-gen-*.yaml' in folder: %s\n", folder)
	}
	matches = append(matches, matches4...)

	for _, match := range matches {
		c, cerr := readConfigFile(match)
		if cerr != nil {
			return nil, errors.Wrapf(cerr, "Error in gen.CreateFromFolder with 'geb-gen.yml' file in folder: %s\n", folder)
		}

		g.Config.Merge(c)
	}

	d, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "designs"))
	if err != nil {
		logger.Info("no 'designs' folder found in generator")
		// return nil, errors.Wrapf(err, "while reading 'designs' folder in: %s\n", folder)
	} else {
		g.Designs = d
	}

	n, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "new"))
	if err != nil {
		logger.Info("no 'new' folder found in generator")
		// return nil, errors.Wrapf(err, "while reading 'new' folder in: %s\n", folder)
	} else {
		g.NewTemplates = n
	}

	p, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "partials"))
	if err != nil {
		logger.Info("no 'partials' folder found in generator")
		// return nil, errors.Wrapf(err, "while reading 'partials' folder in: %s\n", folder)
	} else {
		g.Partials = p
	}

	// Templates is the only required folder in a generator
	r, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "templates"))
	if err != nil {
		return nil, errors.Wrapf(err, "while reading 'repeated' folder in: %s\n", folder)
	}
	g.Templates = r

	logger.Info("!!!!!!!!!!FOUND in generator", "g", g, "config", *g.Config)
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
