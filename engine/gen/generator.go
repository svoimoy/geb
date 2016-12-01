package gen

import (
	"github.com/hofstadter-io/geb/engine/templates"
	"path/filepath"
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

	c, err := ReadConfigFile(filepath.Join(folder, "geb-gen.yml"))
	if err != nil {
		c, err = ReadConfigFile(filepath.Join(folder, "geb-gen.yaml"))
		if err != nil {
			return nil, err
		}
	}
	g.Config = c
	g.SourcePath = folder

	p, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "partials"))
	if err != nil {
		return nil, err
	}
	g.Partials = p

	t, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "templates"))
	if err != nil {
		return nil, err
	}
	g.Templates = t

	r, err := templates.CreateTemplateMapFromFolder(filepath.Join(folder, "repeated"))
	if err != nil {
		return nil, err
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
