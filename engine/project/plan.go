package project

import (
	"path/filepath"

	"github.com/aymerick/raymond"
)

type FileGenData struct {
	Dsl      string
	Gen      string
	File     string
	Data     interface{}
	Template *raymond.Template
	Outfile  string
}

func (P *Project) Plan() error {
	logger.Info("Planning Project")

	plans := []FileGenData{}
	for d_key, D := range P.DslMap {
		logger.Info("    dsl: "+D.Name, "key", d_key)
		for g_key, G := range D.Generators {
			logger.Info("      gen: "+g_key, "gen_cfg", G.Config)
			for t_key, T := range G.Templates {
				fgd := FileGenData{
					Dsl:      d_key,
					Gen:      g_key,
					File:     t_key,
					Template: (*raymond.Template)(T),
					Data:     P.Design,
					Outfile:  filepath.Join(P.Config.OutputDir, d_key, g_key, t_key),
				}
				logger.Info("        file: "+t_key, "fgd", fgd)
				plans = append(plans, fgd)
			}

		}

	}

	P.Plans = plans

	return nil
}
