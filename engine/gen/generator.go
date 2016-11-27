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

	c, err := ReadConfigFile(filepath.Join(folder, "geb.yaml"))
	if err != nil {
		return nil, err
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
