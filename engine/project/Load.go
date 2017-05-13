package project

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryanuber/go-glob"

	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/hofstadter-io/geb/engine/utils"
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
func (P *Project) LoadGenerators() (err error) {
	// HOFSTADTER_START LoadGenerators
	logger.Info("Loading Generators")

	// search for available generators
	err = P.FindAvailableGenerators(nil)
	if err != nil {
		return errors.Wrap(err, "while LoadingGenerators\n")
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
		return errors.New("Did not find DSL in available list: " + s_dsl)
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

		if found {
			logger.Info("    importing", "dsl", s_dsl, "generator", s_gen)

			// for each DSL lookup path
			for _, path := range dslLookupPaths {

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

				// load the dsl (this looks to be redundent if there is more than one generator in the list)
				// except for the overloading aspect...
				dsl_path := filepath.Join(path, s_dsl)
				D, err := dsl.CreateFromFolder(dsl_path)
				if err != nil {
					return err
				}

				// load the generator
				gen_path := filepath.Join(dsl_path, s_gen)
				G, err := gen.CreateFromFolder(gen_path)
				if err != nil {
					return err
				}

				// possibly override the outpur directory
				if gp.OutputDir != "" {
					G.Config.OutputDir = gp.OutputDir
				}

				// add the generator to the dsl
				D.Generators[s_gen] = G

				// possibly merge the dsl/generator into the existing
				orig, ok := P.DslMap[s_dsl]
				logger.Debug("    ", "path", path, "s_dsl", s_dsl, "ok", ok)
				if ok {
					orig.MergeSkipExisting(D)
				} else {
					P.DslMap[s_dsl] = D
				}

				// load any dependent generators here
				for _, depGen := range G.Config.Dependencies.Generators {
					// fmt.Println("Dependent GEN: ", depGen)
					derr := P.LoadGenerator(depGen, dslLookupPaths)
					if derr != nil {
						return errors.Wrap(derr, fmt.Sprintf("while loading dependent generator: %+v\n", depGen))
					}
				}
			}
		}

	} // end loop over dsl generators

	// HOFSTADTER_END   LoadGenerator
	return
}

// HOFSTADTER_BELOW
