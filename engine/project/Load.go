package project

import (
	// HOFSTADTER_START import
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/ryanuber/go-glob"

	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import
)

/*
Name:      Load
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
func (P *Project) LoadGenerators() (err error) {
	// HOFSTADTER_START LoadGenerators
	logger.Info("Loading Generators")

	// search for available generators
	err = P.FindAvailableGenerators(nil)
	if err != nil {
		return errors.Wrap(err, "while FindAvailableGenerators\n")
	}
	logger.Info("  Available:", "avail", P.Available)

	cfg := P.Config.DslConfig
	// loop over the project generators
	for _, gConfig := range cfg.Default {
		err = P.LoadGenerator(gConfig, cfg.Paths)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("while LoadingGenerator: %+v\n", gConfig))
		}
	}

	return nil
	// HOFSTADTER_END   LoadGenerators
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) LoadGenerator(generator gen.GeneratorConfig, dslLookupPaths []string) (err error) {
	// HOFSTADTER_START LoadGenerator
	gp := generator
	s_dsl := gp.Dsl

	// check to see if the dsl is available
	d_dsl, ok := P.Available[s_dsl]
	if !ok {
		return errors.New("Did not find DSL in available list: " + s_dsl + " in " + fmt.Sprint(P.Available))
	}

	// if there are no generators specified, load all
	// how useful is this? does it cause more difficulty than it should?
	if len(gp.Gen) == 0 {
		logger.Info("  importing all for " + gp.Dsl)
		for key, path := range d_dsl.AvailableGenerators {
			logger.Info("    appending", "key", key, "dsl", s_dsl, "generator", path)
			gp.Gen = append(gp.Gen, path)
		}
	}

	// for all of the generators under the dsl in the config file...
	for _, s_gen := range gp.Gen {

		spath := s_gen + "*"

		// check that the generator exists
		found := false
		for _, path := range d_dsl.AvailableGenerators {
			found = glob.Glob(spath, path)
			logger.Debug("GLOB:", "spath", spath, "path", path, "found", found)
			if found {
				break
			}
		} // end for loop looking for gen in available generators

		var lastErr error

		if !found {
			logger.Error("Did not find generator", "gen", spath)
			return errors.Errorf("Did not find generator %q an any dsl paths %v", spath, dslLookupPaths)
		} else {
			logger.Info("    importing", "dsl", s_dsl, "generator", s_gen)

			// for each DSL lookup path
			loaded := false
			for _, path := range dslLookupPaths {

				// Resolve the path for EnvVars, symlinks, existance
				t_path, err := utils.ResolvePath(path)
				// skip it if the file does not exist
				if err != nil {
					if _, ok := err.(*os.PathError); ok {
						lastErr = err
						continue
					}
					if strings.Contains(err.Error(), "no such file or directory") {
						lastErr = err
						continue
					}

					// otherwise return the error
					return errors.Wrapf(err, "in project.LoadGeneratorList")
				}
				path = t_path

				// load the dsl (this looks to be redundent if there is more than one generator in the list)
				// except for the overloading aspect...
				dsl_path := filepath.Join(path, s_dsl)
				D, err := dsl.CreateFromFolder(dsl_path)
				if err != nil {
					lastErr = err
					continue
				}

				// load the generator
				gen_path := filepath.Join(dsl_path, s_gen)
				G, err := gen.CreateFromFolder(gen_path)
				logger.Debug("    ", "path", path, "gen", G, "err", err)
				if err != nil {
					lastErr = err
					continue
				}

				// possibly override the outpur directory
				if gp.OutputDir != "" {
					G.Config.OutputDir = gp.OutputDir
				}

				// add the generator to the dsl
				D.Generators[s_gen] = G
				logger.Debug("    ", "path", path, "dsl", D, "err", err)

				// possibly merge the dsl/generator into the existing
				orig, ok := P.DslMap[s_dsl]
				logger.Debug("    ", "path", path, "s_dsl", s_dsl, "ok", ok)
				if ok {
					// check if the generator has been loaded already
					_, ok := orig.Generators[s_gen]
					if !ok {
						orig.MergeSkipExisting(D)
					} else {
						// otherwise it has already been loaded
						loaded = true
						continue
					}
				} else {
					P.DslMap[s_dsl] = D
				}

				// load any dependent generators here
				for _, depGen := range G.Config.Dependencies.Generators {
					// fmt.Println("Dependent GEN: ", depGen)
					derr := P.LoadGenerator(depGen, dslLookupPaths)
					if derr != nil {
						return errors.Wrap(derr, fmt.Sprintf("while loading dependent generator: %+v %+v\n", s_gen, depGen))
					}
				}

				loaded = true
			}

			if !loaded {
				return errors.Wrap(lastErr, fmt.Sprintf("while loading generator: %+v\n", gp))
			}
		}

	} // end loop over dsl generators

	logger.Debug("Project after gen load", "P", P)

	// HOFSTADTER_END   LoadGenerator
	return
}

// HOFSTADTER_BELOW
