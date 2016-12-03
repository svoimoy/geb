package gen

import (
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/hofstadter-io/geb/engine/templates"
)

type Generator struct {
	SourcePath string
	Config     *Config

	Templates templates.TemplateMap
	Repeated  templates.TemplateMap
	Partials  templates.TemplateMap
}

func NewGenerator() *Generator {
	return &Generator{
		Config:    NewConfig(),
		Templates: templates.NewTemplateMap(),
		Repeated:  templates.NewTemplateMap(),
		Partials:  templates.NewTemplateMap(),
	}
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

	t, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "templates"))
	if err != nil {
		return nil, errors.Wrapf(err, "while reading 'templates' folder in: %s\n", folder)
	}
	g.Templates = t

	r, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "repeated"))
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

func (G *Generator) MergeSkipExisting(stale *Generator) {
	logger.Info("Merging GEN", "existing", G.SourcePath, "stale", stale.SourcePath)
	for path, T := range stale.Templates {
		_, ok := G.Templates[path]
		if ok {
			logger.Info("Skipping template", "template", path)
		} else {
			logger.Info("Adding template", "template", path)
			G.Templates[path] = T
		}
	}
	for path, T := range stale.Repeated {
		_, ok := G.Repeated[path]
		if ok {
			logger.Info("Skipping repeated", "repeated", path)
		} else {
			logger.Info("Adding repeated", "repeated", path)
			G.Repeated[path] = T
		}
	}
	for path, P := range stale.Partials {
		_, ok := G.Partials[path]
		if ok {
			logger.Info("Skipping partial", "partial", path)
		} else {
			logger.Info("Adding partial", "partial", path)
			G.Partials[path] = P
		}
	}
}
