package plan

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"

	"github.ibm.com/hofstadter-io/dotpath"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	// HOFSTADTER_END   import
)

func make_package(dsl_ctx interface{}, dsl_map map[string]*dsl.Dsl, design_data map[string]interface{}) ([]Plan, error) {

	logger.Info("Making Pkg plan")
	logger.Debug("  context", "dsl_ctx", dsl_ctx)

	// get the ctx path for later comparison against dsl
	ictx_path, err := dotpath.Get("ctx_path", dsl_ctx, true)
	if err != nil {
		return nil, errors.New("ctx_path not found, in make_type")
	}
	ctx_path, ok := ictx_path.(string)
	if !ok {
		return nil, errors.New("ctx_path is not a string, in make_type")
	}

	ctx_flds := strings.Split(ctx_path, ".")
	ctx_dir := ""
	if len(ctx_flds) > 2 {
		ctx_dir = filepath.Join(ctx_flds[1 : len(ctx_flds)-1]...)
	}
	ctx_dsl := ctx_flds[0]

	plans := []Plan{}

	logger.Debug("Making Pkg plan", "dsl_map", dsl_map, "ctx_dsl", ctx_dsl)

	// Loop over DSLs in the plans
	for d_key, D := range dsl_map {
		// ... comparing the dsl type to the design type
		if d_key != ctx_dsl {
			continue
		}
		logger.Info("    dsl: "+D.Config.Name, "d_key", d_key, "ctx_dsl", ctx_dsl, "ctx_path", ctx_path)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Info("      gen: "+g_key, "gen_cfg", G.Config)

			G_key := filepath.Join(d_key, g_key)
			if G.Config.OutputDir != "" {
				G_key = G.Config.OutputDir
			}

			//
			//  NORMAL TEMPLATES
			//
			// Render the normal templates
			for t_key, T := range G.Templates {
				t_ray := (*raymond.Template)(T)
				outfile := filepath.Join(ctx_dir, G_key, t_key)

				// build up the plan data struct
				p := Plan{
					Dsl:        d_key,
					Gen:        g_key,
					File:       t_key,
					Template:   t_ray,
					Data:       design_data,
					Outfile:    outfile,
					DslContext: dsl_ctx,
				}
				logger.Info("        template file: "+t_key, "plan", p)

				// add the plan to a linear list to be rendered
				plans = append(plans, p)
			} // End of normal template processing

			//
			//  REPEAT TEMPLATES
			//
			// Start of repeat processing section:
			repeats := G.Config.Repeated
			if len(repeats) == 0 {
				logger.Debug("       skipping pkg repeat: "+D.Config.Type, "name", D.Config.Name, "repeats", repeats)
				continue
			}
			logger.Info("Repeated found in config:", "count", len(repeats), "repeats", repeats)
			logger.Info("      doing pkg repeat: "+D.Config.Type, "name", D.Config.Name, "d_key", d_key)

			// Render the repeated templates
			for _, R := range repeats {
				logger.Info("Processing Repeated Field: '"+R.Name+"'", "field", R.Field, "dsl_ctx", dsl_ctx)

				repeat_elems, err := dotpath.Get(R.Field, dsl_ctx, false)
				if err != nil || repeat_elems == nil {
					logger.Debug("Skipping Repeated Field: '"+R.Name+"'", "err", err, "repeat_elems", repeat_elems)

					continue
				}

				logger.Debug("Doing Repeated Field: '" + R.Name + "'")
				var c_slice []interface{}

				switch M := repeat_elems.(type) {

				case map[string]interface{}:
					c_slice = append(c_slice, M)

				case map[interface{}]interface{}:
					c_slice = append(c_slice, M)

				case []interface{}:
					for _, elem := range M {
						c_slice = append(c_slice, elem)
					}

				default:
					logger.Info("input is not a map or slice", "input", M)

				}

				logger.Info("Done adding to c_slice", "c_slice", c_slice)

				/*
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
				*/

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

						outfile := filepath.Join(ctx_dir, G_key, OF_name)

						// build up the plan data struct
						fgd := Plan{
							Dsl:      d_key,
							Gen:      g_key,
							File:     t_key,
							Template: t_ray,
							Data:     design_data,
							Outfile:  outfile,

							DslContext:      dsl_ctx,
							RepeatedContext: local_ctx,
						}
						// logger.Info("        planned repeat file: "+t_key, "index", idx)
						// logger.Debug("          data...", "fgd", fgd, "index", idx)

						// add the plan to a linear list to be rendered
						plans = append(plans, fgd)

					} // END of context loop 'c_slice'
					logger.Debug("    end repeat loop: ", "repeat", R.Name, "in", t_key, "c_slice", c_slice)

				}

			} // End of repeated template processing

		} // End Generator loop

	} // End DSL loop

	return plans, nil
}
