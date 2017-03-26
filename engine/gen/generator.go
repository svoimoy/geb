package gen

import (
	// HOFSTADTER_START import
	"path/filepath"

	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/templates"
	// HOFSTADTER_END   import
)

// Name:      generator
// Namespace: engine.gen
// Version:   0.0.1

type Generator struct {
	Config     *Config               `json:"config" xml:"config" yaml:"config" form:"config" query:"config" `
	SourcePath string                `json:"source-path" xml:"source-path" yaml:"source-path" form:"source-path" query:"source-path" `
	Templates  templates.TemplateMap `json:"templates" xml:"templates" yaml:"templates" form:"templates" query:"templates" `
	Repeated   templates.TemplateMap `json:"repeated" xml:"repeated" yaml:"repeated" form:"repeated" query:"repeated" `
	Partials   templates.TemplateMap `json:"partials" xml:"partials" yaml:"partials" form:"partials" query:"partials" `
}

func NewGenerator() *Generator {
	return &Generator{

		Config: NewConfig(),
	}
	// loop over fields looking for pointers
}

// HOFSTADTER_BELOW

func New() *Generator {
	return NewGenerator()
}

func CreateFromFolder(folder string) (*Generator, error) {
	g := NewGenerator()

	c, cerr := ReadConfigFile(filepath.Join(folder, "geb-gen.yml"))
	if cerr != nil {
		cerr = errors.Wrapf(cerr, "Error in gen.CreateFromFolder with 'geb-gen.yml' file in folder: %s\n", folder)
		c, cerr = ReadConfigFile(filepath.Join(folder, "geb-gen.yaml"))
		if cerr != nil {
			return nil, errors.Wrapf(cerr, "Error in gen.CreateFromFolder with 'geb-gen.yaml' file in folder: %s\n", folder)
		}
	}

	g.Config = c
	g.SourcePath = folder

	p, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "partials"))
	if err != nil {
		return nil, errors.Wrapf(err, "while reading 'partials' folder in: %s\n", folder)
	}
	g.Partials = p

	/*
		t, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "templates"))
		if err != nil {
			return nil, errors.Wrapf(err, "while reading 'templates' folder in: %s\n", folder)
		}
		g.Templates = t
	*/

	r, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "templates"))
	if err != nil {
		return nil, errors.Wrapf(err, "while reading 'repeated' folder in: %s\n", folder)
	}
	g.Repeated = r

	return g, nil
}

func (G *Generator) MergeOverwrite(fresh *Generator) {
	logger.Info("Merging GEN", "existing", G.SourcePath, "fresh", fresh.SourcePath)
	for path, T := range fresh.Templates {
		_, ok := G.Templates[path]
		if ok {
			logger.Info("Overriding template", "template", path)
		} else {
			logger.Info("Adding template", "template", path)
		}
		G.Templates[path] = T
	}
	for path, T := range fresh.Repeated {
		_, ok := G.Repeated[path]
		if ok {
			logger.Info("Overriding repeated", "repeated", path)
		} else {
			logger.Info("Adding repeated", "repeated", path)
		}
		G.Repeated[path] = T
	}
	for path, P := range fresh.Partials {
		_, ok := G.Partials[path]
		if ok {
			logger.Info("Overriding partial", "partial", path)
		} else {
			logger.Info("Adding partial", "partial", path)
		}
		G.Partials[path] = P
	}
}

func (G *Generator) MergeSkipExisting(fresh *Generator) {
	logger.Info("Merging GEN", "existing", G.SourcePath, "fresh", fresh.SourcePath)
	for path, T := range fresh.Templates {
		_, ok := G.Templates[path]
		if ok {
			logger.Info("Skipping template", "template", path)
		} else {
			logger.Info("Adding template", "template", path)
			G.Templates[path] = T
		}
	}
	for path, T := range fresh.Repeated {
		_, ok := G.Repeated[path]
		if ok {
			logger.Info("Skipping repeated", "repeated", path)
		} else {
			logger.Info("Adding repeated", "repeated", path)
			G.Repeated[path] = T
		}
	}
	for path, P := range fresh.Partials {
		_, ok := G.Partials[path]
		if ok {
			logger.Info("Skipping partial", "partial", path)
		} else {
			logger.Info("Adding partial", "partial", path)
			G.Partials[path] = P
		}
	}
}
