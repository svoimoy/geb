package project

import (
	"errors"

	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/ryanuber/go-glob"
)

type Project struct {
	// Read from project directories
	Config *Config
	Design *design.Design
	// LocalTemplates

	// Generators []Generator
	// Pipelines  []Pipeline
}

func NewProject() *Project {
	return &Project{}
}

func (P *Project) Load(filename string) error {

	logger.Info("Reading config file", "filename", filename)
	c, err := ReadConfigFile(filename)
	if err != nil {
		logger.Crit("While reading project config", "filename", filename, "error", err)
		return err
	}
	P.Config = c

	d_dir := P.Config.DesignDir
	logger.Info("Reading designs", "folder", d_dir)
	d, err := design.CreateFromFolder(d_dir)
	if err != nil {
		logger.Crit("While reading project designs", "folder", d_dir, "error", err)
		return err
	}
	P.Design = d

	err = P.LoadGenerators()
	if err != nil {
		logger.Crit("While loading generators", "error", err)
		return err
	}

	return nil

}

func (P *Project) LoadGenerators() error {

	logger.Info("Loading Generators")
	cfg := P.Config.DslConfig

	logger.Info("DSL override order (first to last):")
	for _, path := range cfg.Paths {
		logger.Info("  ", "path", path)
	}
	dsl_dir := "dsl"
	available_dsls, err := dsl.FindAvailable(dsl_dir)
	if err != nil {
		return err
	}

	logger.Info("Importing generators:")
	for _, gp := range cfg.Default {
		s_dsl := gp.Dsl

		d_dsl, ok := available_dsls[s_dsl]
		if !ok {
			return errors.New("Unknown DSL: " + s_dsl)
		}

		if len(gp.Gen) == 0 {
			for _, path := range d_dsl.AvailableGenerators {
				logger.Info("    importing", "dsl", s_dsl, "generator", path)
			}

		} else {
			for _, s_gen := range gp.Gen {

				spath := s_gen + "*"

				found := false
				for _, path := range d_dsl.AvailableGenerators {
					found = glob.Glob(spath, path)
					logger.Debug("GLOB:", "spath", spath, "path", path, "found", found)
					if found {
						break
					}
				} // end for loop looking for gen in available generators

				if found {
					logger.Info("    importing", "dsl", s_dsl, "generator", s_gen)
					gen.NewGenerator()
				}

			} // end loop over dsl generators
		}

	} // end loop over default dsls

	return nil
}
