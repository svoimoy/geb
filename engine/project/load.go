package project

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/ryanuber/go-glob"
)

func (P *Project) Load(filename string, generators []string) error {

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

	err = P.LoadGenerators(generators)
	if err != nil {
		logger.Crit("While loading generators", "error", err)
		return err
	}

	return nil

}

func (P *Project) LoadGenerators(generators []string) error {

	logger.Info("Loading Generators")
	cfg := P.Config.DslConfig

	logger.Info("DSL override order (first to last):")
	available_dsls := map[string]*dsl.Dsl{}
	for _, path := range cfg.Paths {
		if path[:2] == "~/" {
			usr, _ := user.Current()
			home := usr.HomeDir
			path = strings.Replace(path, "~", home, 1)
		}

		// skip it?
		_, err := os.Lstat(path)
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				continue
			}
			return err
		}

		// find whats available
		avail, err := dsl.FindAvailable(path)
		if err != nil {
			return err
		}
		for key, val := range avail {
			existing, ok := available_dsls[key]
			if ok {
				existing.MergeAvailable(val)
				available_dsls[key] = existing
			} else {
				available_dsls[key] = val
			}
		}
	}

	if len(generators) == 0 {
		return P.LoadDefaultGenerators(available_dsls)
	} else {
		return P.LoadGeneratorList(available_dsls, generators)
	}
}

func (P *Project) LoadDefaultGenerators(available_dsls map[string]*dsl.Dsl) error {
	logger.Info("Importing generators:")
	cfg := P.Config.DslConfig
	for _, gp := range cfg.Default {
		s_dsl := gp.Dsl

		d_dsl, ok := available_dsls[s_dsl]
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
					if path[:2] == "~/" {
						usr, _ := user.Current()
						home := usr.HomeDir
						path = strings.Replace(path, "~", home, 1)
					}

					// skip it?
					info, err := os.Lstat(path)
					if err != nil {
						if _, ok := err.(*os.PathError); ok {
							continue
						}
						return err
					}

					if info.Mode()&os.ModeSymlink != 0 {
						dir, err := os.Readlink(path)
						if err != nil {
							return err
						}
						path = dir
					}

					dsl_path := filepath.Join(path, s_dsl)
					D, err := dsl.LoadDsl(dsl_path)
					if err != nil {
						return err
					}

					gen_path := filepath.Join(dsl_path, s_gen)
					G, err := gen.CreateFromFolder(gen_path)
					if err != nil {
						return err
					}
					D.Generators[s_gen] = G

					orig, ok := P.DslMap[s_dsl]
					logger.Debug("    ", "path", path, "s_dsl", s_dsl, "ok", ok)
					if ok {
						orig.MergeOverwrite(D)
					} else {
						P.DslMap[s_dsl] = D
					}
				}
			}

		} // end loop over dsl generators

	} // end loop over default dsls

	return nil
}

func (P *Project) LoadGeneratorList(available_dsls map[string]*dsl.Dsl, generators []string) error {
	logger.Info("Importing generator list:")
	logger.Info("  Available:", "avail", available_dsls)
	logger.Info("  Generators:", "gener", generators)

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
				_, ok := available_dsls[l_dsl]
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
		d_dsl, ok := available_dsls[s_dsl]
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
					if path[:2] == "~/" {
						usr, _ := user.Current()
						home := usr.HomeDir
						path = strings.Replace(path, "~", home, 1)
					}

					// skip it?
					info, err := os.Lstat(path)
					if err != nil {
						if _, ok := err.(*os.PathError); ok {
							continue
						}
						return err
					}

					if info.Mode()&os.ModeSymlink != 0 {
						dir, err := os.Readlink(path)
						if err != nil {
							return err
						}
						path = dir
					}

					dsl_path := filepath.Join(path, s_dsl)
					D, err := dsl.LoadDsl(dsl_path)
					if err != nil {
						return err
					}

					logger.Debug("  gen path;", "dsl_path", dsl_path, "gpath", gpath)
					gen_path := filepath.Join(dsl_path, gpath)
					G, err := gen.CreateFromFolder(gen_path)
					if err != nil {
						return err
					}
					D.Generators[s_gen] = G

					orig, ok := P.DslMap[s_dsl]
					logger.Debug("    ", "path", path, "s_dsl", s_dsl, "ok", ok)
					if ok {
						orig.MergeOverwrite(D)
					} else {
						P.DslMap[s_dsl] = D
					}
				}
			}
		} // end for loop looking for gen in available generators
	} // End loop over generators

	return nil
}
