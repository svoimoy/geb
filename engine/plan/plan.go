package plan

import (
	"github.com/pkg/errors"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/templates"
	"github.com/hofstadter-io/geb/engine/utils"
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

				// build up the plan data struct
				p := Plan{
					Dsl:      d_key,
					Gen:      g_key,
					File:     t_key,
					Template: t_ray,
					Data:     design_data,
					Outfile:  filepath.Join(d_key, g_key, t_key),
				}
				logger.Info("        template file: "+t_key, "plan", p)

				// add the plan to a linear list to be rendered
				plans = append(plans, p)
			} // End of normal template processing

			// Start of repeat processing section:
			//
			// only do repeats for actual dsls and
			//   when there are repeats
			//
			repeats := G.Config.Repeated
			if D.Config.Type != "dsl" || len(repeats) == 0 {
				logger.Debug("       skipping dsl repeat: "+D.Config.Type, "name", D.Config.Name)
				continue
			}
			logger.Info("Repeated found in config:", "count", len(repeats), "repeats", repeats)
			logger.Info("      doing dsl repeat: "+D.Config.Type, "name", D.Config.Name, "d_key", d_key)

			// Render the repeated templates
			// Get the root of the data to index into
			for k, _ := range dsl_map {
				logger.Debug("       - dsl keys", "key", k)
			}
			data, ok := design_data["dsl"].(map[string]interface{})[d_key]
			if !ok {
				logger.Error("Did not find DSL data", "d_key", d_key)
				return nil, errors.New("Unknown dsl design: " + d_key)
			}

			for _, R := range repeats {
				logger.Info("Processing Repeated Field: '" + R.Name + "'")

				// look up field
				collection, err := utils.GetByPath(R.Field, data)
				if err != nil {
					return nil, errors.Wrapf(err, "looking up by path:  repeat(%s)  path(%s) in data:\n%+v\n\n", R.Name, R.Field, data)
				}

				c_slice, ok := collection.([]interface{})
				if !ok {
					return nil, errors.New("Collection is not a list: " + R.Field)
				}

				logger.Info("   Collection count", "collection", R.Field, "count", len(c_slice))
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

						templates.AddHelpers(tpl)
						OF_name, err := tpl.Exec(val)
						if err != nil {
							return nil, errors.Wrap(err, "in MakePlans\n")
						}

						OF_name = strings.ToLower(OF_name)
						logger.Info("OFNAME", "name", OF_name)
						outfile := filepath.Join(d_key, g_key, OF_name)

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
						logger.Info("        planned repeat file: "+t_key, "fgd", fgd, "index", idx)

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
