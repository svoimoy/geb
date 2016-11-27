package gen

import "path/filepath"

type Generator struct {
	SourcePath string
	Config     *Config

	Templates TemplateMap
	Partials  TemplateMap
}

func NewGenerator() *Generator {
	return &Generator{
		Config:    NewConfig(),
		Templates: NewTemplateMap(),
		Partials:  NewTemplateMap(),
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

	t, err := CreateTemplateMapFromFolder(filepath.Join(folder, "templates"))
	if err != nil {
		return nil, err
	}
	g.Templates = t

	p, err := CreateTemplateMapFromFolder(filepath.Join(folder, "partials"))
	if err != nil {
		return nil, err
	}
	g.Partials = p

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
