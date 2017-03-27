package project

// package publicFiles

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryanuber/go-glob"

	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/gen"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import
)

/*
Name:      load
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
func (P *Project) LoadDefaultGenerators() (err error) {
	// HOFSTADTER_START LoadDefaultGenerators
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
	// HOFSTADTER_END   LoadDefaultGenerators
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) LoadGenerators(generators []string) (err error) {
	// HOFSTADTER_START LoadGenerators
	err = P.FindAvailableGenerators(nil)
	if err != nil {
		return errors.Wrap(err, "while LoadingGenerators\n")
	}
	if len(generators) == 0 {
		return P.LoadDefaultGenerators()
	} else {
		return P.LoadGeneratorList(generators)
	}
	// HOFSTADTER_END   LoadGenerators
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) LoadGeneratorList(generators []string) (err error) {
	// HOFSTADTER_START LoadGeneratorList
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
	// HOFSTADTER_END   LoadGeneratorList
	return
}

// HOFSTADTER_BELOW
