package project

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryanuber/go-glob"
	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/gen"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
)

func (P *Project) Load(filename string, generators []string) error {

	logger.Info("Reading config file", "filename", filename)
	c, err := ReadConfigFile(filename)
	if err != nil {
		return errors.Wrap(err, "while reading project config file: "+filename)
	}
	P.Config = c
	logger.Debug("Project Config", "config", P.Config)

	err = P.LoadGenerators(generators)
	if err != nil {
		return errors.Wrap(err, "while loading generators\n")
	}

	d_dir := P.Config.DesignDir
	logger.Info("Reading designs", "folder", d_dir)
	d, err := design.CreateFromFolder(d_dir)
	if err != nil {
		return errors.Wrapf(err, "While reading design folder: %s\n", d_dir)
	}
	P.Design = d
	logger.Debug("Project Design", "design", P.Design)

	return nil

}

// This function seraches for available dsls and generators.
// If no paths are provided, it uses the project defaults, or geb defaults.
// This function may be called repeatedly to add and merge.
func (P *Project) FindAvailableGenerators(paths []string) error {
	logger.Info("Searching for Generators")

	// If no paths are provided, use those defined in the configuration
	if len(paths) == 0 {
		paths = P.Config.DslConfig.Paths
	}

	logger.Info("DSL override order (first to last):", "paths", paths)
	if P.Available == nil {
		P.Available = map[string]*dsl.Dsl{}
	}
	for _, path := range paths {

		// Resolve the path for EnvVars, symlinks, existance
		t_path, err := utils.ResolvePath(path)
		// skip it if the file does not exist
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				continue
			}
			if strings.Contains(err.Error(), "no such file or directory") {
				continue
			}

			// otherwise return the error
			return errors.Wrapf(err, "in project.FindAvailGens\n")
		}
		path = t_path

		// Find out what's available
		avail, err := dsl.FindAvailable(path)
		if err != nil {
			return errors.Wrapf(err, "in proj.FindAvailGens %v\n", paths)
		}
		for key, val := range avail {
			existing, ok := P.Available[key]
			if ok {
				existing.MergeAvailable(val)
				P.Available[key] = existing
			} else {
				P.Available[key] = val
			}
		}
	}

	return nil
}

func (P *Project) LoadGenerators(generators []string) error {
	err := P.FindAvailableGenerators(nil)
	if err != nil {
		return errors.Wrap(err, "while LoadingGenerators\n")
	}
	if len(generators) == 0 {
		return P.LoadDefaultGenerators()
	} else {
		return P.LoadGeneratorList(generators)
	}
}

func (P *Project) LoadDefaultGenerators() error {
	logger.Info("Loading Default Generators")
	logger.Info("  Available:", "avail", P.Available)

	cfg := P.Config.DslConfig
	for _, gp := range cfg.Default {
		s_dsl := gp.Dsl

		d_dsl, ok := P.Available[s_dsl]
		if !ok {
			return errors.New("Did not find DSL in available list: " + s_dsl)
		}

		if len(gp.Gen) == 0 {
			logger.Info("  importing all for " + gp.Dsl)
			for key, path := range d_dsl.AvailableGenerators {
				logger.Info("    appending", "key", key, "dsl", s_dsl, "generator", path)
				gp.Gen = append(gp.Gen, path)
			}
		}

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
				for _, path := range cfg.Paths {

					// Resolve the path for EnvVars, symlinks, existance
					t_path, err := utils.ResolvePath(path)
					// skip it if the file does not exist
					if err != nil {
						if _, ok := err.(*os.PathError); ok {
							continue
						}
						if strings.Contains(err.Error(), "no such file or directory") {
							continue
						}

						// otherwise return the error
						return errors.Wrapf(err, "in project.LoadGeneratorList")
					}
					path = t_path

					dsl_path := filepath.Join(path, s_dsl)
					D, err := dsl.CreateFromFolder(dsl_path)
					if err != nil {
						return err
					}

					gen_path := filepath.Join(dsl_path, s_gen)
					G, err := gen.CreateFromFolder(gen_path)
					if err != nil {
						return err
					}

					if gp.OutputDir != "" {
						G.Config.OutputDir = gp.OutputDir
					}
					D.Generators[s_gen] = G

					orig, ok := P.DslMap[s_dsl]
					logger.Debug("    ", "path", path, "s_dsl", s_dsl, "ok", ok)
					if ok {
						orig.MergeSkipExisting(D)
					} else {
						P.DslMap[s_dsl] = D
					}
				}
			}

		} // end loop over dsl generators

	} // end loop over default dsls

	return nil
}

// This function loads a list of generators into the project
// It looks through the list of paths defined in the project configuration
func (P *Project) LoadGeneratorList(generators []string) error {
	logger.Info("Loading Custom Generators", generators, generators)
	logger.Info("  Available:", "avail", P.Available)

	for _, g_str := range generators {
		fields := strings.Split(g_str, "/")
		logger.Info("  - "+g_str, "fields", fields)
		s_dsl := ""
		s_gen := ""
		if len(fields) == 1 {
			s_dsl = fields[0]
			logger.Info("    found: " + s_dsl + " *")
		} else {
			for i := len(fields) - 1; i > 0; i-- {
				l_dsl := strings.Join(fields[:i], "/")
				_, ok := P.Available[l_dsl]
				if ok {
					s_dsl = l_dsl
					s_gen = strings.Join(fields[i:], "/")
					logger.Info("    found: " + s_dsl + " " + s_gen)
					break
				}
			}
		}

		if s_dsl == "" {
			return errors.New("Did not find DSL in available list: " + s_dsl)
		}
		d_dsl, ok := P.Available[s_dsl]
		if !ok {
			return errors.New("Really did not find DSL in available list: " + s_dsl)
		}

		spath := s_gen + "*"
		gpath := ""

		found := false
		for _, path := range d_dsl.AvailableGenerators {
			found = glob.Glob(spath, path)
			logger.Debug("GLOB:", "spath", spath, "path", path, "found", found)
			if found {
				gpath = path
				logger.Warn("    importing", "dsl", s_dsl, "generator", s_gen, "spath", spath)
				for _, path := range P.Config.DslConfig.Paths {

					t_path, err := utils.ResolvePath(path)
					// skip it if the file does not exist
					if err != nil {
						if _, ok := err.(*os.PathError); ok {
							continue
						}
						if strings.Contains(err.Error(), "no such file or directory") {
							continue
						}

						// otherwise return the error
						return errors.Wrapf(err, "in project.LoadGeneratorList")
					}
					path = t_path

					// load the dsl config file
					dsl_path := filepath.Join(path, s_dsl)
					D, err := dsl.CreateFromFolder(dsl_path)
					if err != nil {
						return errors.Wrapf(err, "in project.LoadGeneratorList")
					}

					//
					logger.Debug("  gen path;", "dsl_path", dsl_path, "gpath", gpath)
					gen_path := filepath.Join(dsl_path, gpath)
					G, err := gen.CreateFromFolder(gen_path)
					if err != nil {
						return errors.Wrapf(err, "in project.LoadGeneratorList")
					}
					D.Generators[s_gen] = G

					orig, ok := P.DslMap[s_dsl]
					logger.Debug("    ", "path", path, "s_dsl", s_dsl, "ok", ok)
					if ok {
						orig.MergeSkipExisting(D)
					} else {
						P.DslMap[s_dsl] = D
					}

				}
			}
		} // end for loop looking for gen in available generators
	} // End loop over generators

	return nil
}
