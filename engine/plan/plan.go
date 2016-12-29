package plan

import (
	"github.com/pkg/errors"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"
	"github.ibm.com/hofstadter-io/dotpath"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/templates"
)

type Plan struct {
	Dsl      string
	Gen      string
	File     string
	Data     map[string]interface{}
	Template *raymond.Template
	Outfile  string

	RepeatedContext interface{}
}

func MakePlans(dsl_map map[string]*dsl.Dsl, design_data map[string]interface{}) ([]Plan, error) {
	logger.Info("Planning Project")

	plans := []Plan{}

	logger.Info("Making the plans")

	// Loop over DSLs in the plans
	for d_key, D := range dsl_map {
		logger.Info("    dsl: "+D.Config.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Info("      gen: "+g_key, "gen_cfg", G.Config)

			// Render the normal templates
			for t_key, T := range G.Templates {
				t_ray := (*raymond.Template)(T)

				G_key := filepath.Join(d_key, g_key)
				if G.Config.OutputDir != "" {
					G_key = G.Config.OutputDir
				}
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

			//
			//
			//
			//
			//
			//
			//
			// Start of repeat processing section:
			//
			// only do repeats for actual dsls and
			//   when there are repeats
			//
			repeats := G.Config.Repeated
			if !(D.Config.Type == "dsl" || D.Config.Type == "type") || len(repeats) == 0 {
				logger.Debug("       skipping dsl repeat: "+D.Config.Type, "name", D.Config.Name, "repeats", repeats)
				continue
			}
			logger.Info("Repeated found in config:", "count", len(repeats), "repeats", repeats)
			logger.Info("      doing dsl repeat: "+D.Config.Type, "name", D.Config.Name, "d_key", d_key)

			// Render the repeated templates
			// Get the root of the data to index into
			for k, _ := range dsl_map {
				logger.Debug("       - dsl key: "+k, "key", k)
			}
			var data interface{}
			switch D.Config.Type {
			case "dsl":
				d, ok := design_data["dsl"].(map[string]interface{})[d_key]
				if !ok {
					logger.Error("Did not find DSL data", "d_key", d_key, "design_data", design_data)
					return nil, errors.Errorf("Did not find design data in your project for dsl: " + d_key)
				}
				data = d
			case "type":
				d, ok := design_data["type"].(map[string]interface{})
				if !ok || len(d) == 0 {
					logger.Error("Did not find any Type data", "design_data", design_data)
					return nil, errors.Errorf("Did not find design data in your project for dsl: " + d_key)
				}
				data = d

			}

			for _, R := range repeats {
				logger.Info("Processing Repeated Field: '" + R.Name + "'")

				// look up field
				// lookup := R.Field + "." + strings.ToLower(R.Name)
				// logger.Info("    lookup: "+lookup, "R", R)
				var c_slice []interface{}
				switch D.Config.Type {
				case "dsl":
					collection, err := dotpath.Get(R.Field, data)
					if err != nil {
						return nil, errors.Wrapf(err, "looking up by path:  repeat(%s)  path(%s) in data:\n%+v\n\n", R.Name, R.Field, data)
					}

					tmp_c_slice, ok := collection.([]interface{})
					if !ok {
						logger.Info("Collection not a slice", "collection", collection)
						// return nil, errors.New("Collection is not a list: " + R.Field)
						c_slice = []interface{}{collection}
					} else {
						c_slice = tmp_c_slice
					}

				case "type":
					for _, typ := range design_data["type"].(map[string]interface{}) {
						local_typ := typ
						logger.Crit("Adding typ to c_slice", "typ", local_typ, "types", design_data["type"])

						// Recurse over type map here, looking for elements...
						// which have both name and namespace set.
						// This is so we can have nested directories and packages of types

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
										logger.Error("elem is not a mapSI", "elem", elem)
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
										logger.Error("elem is not a mapII", "elem", elem)
										continue
									}

									if has_name && has_namespace {
										c_slice = append(c_slice, elem)
									} else {
										extract_elems(elem)
									}
								}

							default:
								logger.Error("input is not a map", "input", MAP)

							}

						}

						extract_elems(local_typ)
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
						local_ctx := val
						logger.Debug("     context", "val", local_ctx, "idx", idx)

						of_tpl_source := t_pair.Out
						tpl, err := raymond.Parse(of_tpl_source)
						if err != nil {
							return nil, errors.Wrap(err, "in MakePlans\n")
						}
						logger.Debug("       tpair:", "tpair", t_pair, "val", val)

						templates.AddHelpers(tpl)
						OF_name, err := tpl.Exec(val)
						if err != nil {
							return nil, errors.Wrap(err, "in MakePlans\n")
						}

						OF_name = strings.ToLower(OF_name)
						logger.Info("OFNAME", "name", OF_name)

						G_key := filepath.Join(d_key, g_key)
						if G.Config.OutputDir != "" {
							G_key = G.Config.OutputDir
						}

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
