package plan

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"

	"github.com/aymerick/raymond"
	"github.com/kr/pretty"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	// HOFSTADTER_END   import
)

func make_packages(dsl_map map[string]*dsl.Dsl, design_data map[string]interface{}) ([]Plan, error) {

	plans := []Plan{}

	logger.Info("Making Pkg plans")

	// Loop over DSLs in the plans
	for d_key, D := range dsl_map {
		if !(D.Config.Type == "pkg") {
			continue
		}
		logger.Info("    pkg: "+D.Config.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Info("      gen: "+g_key, "gen_cfg", G.Config)

			G_key := filepath.Join(d_key, g_key)
			if G.Config.OutputDir != "" {
				G_key = G.Config.OutputDir
			}

			// Render the normal templates
			for t_key, T := range G.Templates {
				t_ray := (*raymond.Template)(T)
				outfile := filepath.Join(G_key, t_key)

				// build up the plan data struct
				p := Plan{
					Dsl:      d_key,
					Gen:      g_key,
					File:     t_key,
					Template: t_ray,
					Data:     design_data,
					Outfile:  outfile,
				}
				logger.Info("        template file: "+t_key, "plan", p)

				// add the plan to a linear list to be rendered
				plans = append(plans, p)
			} // End of normal template processing

			// Start of repeat processing section:
			repeats := G.Config.Repeated
			if len(repeats) == 0 {
				logger.Debug("       skipping pkg repeat: "+D.Config.Type, "name", D.Config.Name, "repeats", repeats)
				continue
			}
			logger.Info("Repeated found in config:", "count", len(repeats), "repeats", repeats)
			logger.Info("      doing pkg repeat: "+D.Config.Type, "name", D.Config.Name, "d_key", d_key)

			// Render the repeated templates
			d, ok := design_data["pkg"].(map[string]interface{})
			if !ok || len(d) == 0 {
				logger.Debug("Did not find any Type data", "design_data", fmt.Sprint("\n\n%# v\n\n", pretty.Formatter(design_data)))
				return nil, errors.Errorf("Did not find design data in your project for dsl: " + d_key)
			}

			for _, R := range repeats {
				logger.Info("Processing Repeated Field: '" + R.Name + "'")

				var c_slice []interface{}
				// look up field
				for _, pkg := range design_data["pkg"].(map[string]interface{}) {
					local_pkg := pkg
					logger.Debug("Adding pkg to c_slice", "pkg", local_pkg)

					// Recurse over pkg map here, looking for elements...
					// which have both name and namespace set.
					// This is so we can have nested directories and packages of packages

					var extract_elems func(interface{})

					extract_elems = func(MAP interface{}) {
						switch M := MAP.(type) {

						case map[string]interface{}:
							for _, elem := range M {
								has_name, has_namespace := false, false

								switch E := elem.(type) {
								case map[string]interface{}:
									if _, ok := E["name"]; ok {
										has_name = true
									}
									if _, ok := E["namespace"]; ok {
										has_namespace = true
									}

								case map[interface{}]interface{}:
									if _, ok := E["name"]; ok {
										has_name = true
									}
									if _, ok := E["namespace"]; ok {
										has_namespace = true
									}

								default:
									logger.Debug("elem is not a mapSI", "elem", elem)
									continue
								}

								if has_name && has_namespace {
									c_slice = append(c_slice, elem)
								} else {
									extract_elems(elem)
								}
							}

						case map[interface{}]interface{}:
							for _, elem := range M {
								has_name, has_namespace := false, false

								switch E := elem.(type) {
								case map[string]interface{}:
									if _, ok := E["name"]; ok {
										has_name = true
									}
									if _, ok := E["namespace"]; ok {
										has_namespace = true
									}

								case map[interface{}]interface{}:
									if _, ok := E["name"]; ok {
										has_name = true
									}
									if _, ok := E["namespace"]; ok {
										has_namespace = true
									}

								default:
									logger.Debug("elem is not a mapII", "elem", elem)
									continue
								}

								if has_name && has_namespace {
									c_slice = append(c_slice, elem)
								} else {
									extract_elems(elem)
								}
							}

						}

						extract_elems(local_pkg)
					}
					logger.Info("Done adding to c_slice", "c_slice", c_slice)
				}

				// flattern c_slice
				// ....
				// ....
				tmp_c_slice := []interface{}{}
				for _, elem := range c_slice {
					if A, ok := elem.([]interface{}); ok {
						for _, a := range A {
							tmp_c_slice = append(tmp_c_slice, a)
						}
					} else {
						tmp_c_slice = append(tmp_c_slice, elem)
					}
				}
				c_slice = tmp_c_slice

				logger.Info("   Collection count", "collection", R.Field, "count", len(c_slice), "c_slice", c_slice)
				for _, t_pair := range R.Templates {
					logger.Info("    Looking for repeat template: ", "t_pair", t_pair, "in", G.Repeated)
					t_key := t_pair.In

					T, ok := G.Repeated[t_key]
					if !ok {
						return nil, errors.New("Unknown repeat template: " + t_key)
					}
					t_ray := (*raymond.Template)(T)
					logger.Debug("        found repeat template: ", "repeat", R.Name, "in", t_key)

					for idx, val := range c_slice {
						// needed because of range iteration behavior
						local_ctx := val

						logger.Debug("     context", "val", local_ctx, "idx", idx)

						OF_name, err := determine_outfile_name(t_pair.Out, val)
						if err != nil {
							return nil, errors.Wrap(err, "in make_pkgs\n")
						}
						logger.Info("OFNAME", "name", OF_name)

						outfile := filepath.Join(G_key, OF_name)

						// build up the plan data struct
						fgd := Plan{
							Dsl:      d_key,
							Gen:      g_key,
							File:     t_key,
							Template: t_ray,
							Data:     design_data,
							Outfile:  outfile,

							RepeatedContext: local_ctx,
						}
						// logger.Info("        planned repeat file: "+t_key, "index", idx)
						// logger.Debug("          data...", "fgd", fgd, "index", idx)

						// add the plan to a linear list to be rendered
						plans = append(plans, fgd)

					} // END of context loop 'c_slice'
					logger.Debug("    end repeat loop: ", "repeat", R.Name, "in", t_key, "c_slice", c_slice)

				}

			}
			// End of repeated template processing

		} // End Generator loop

	} // End DSL loop
	return plans, nil
}
