package project

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/hofstadter-io/geb/engine/templates"
	"github.com/hofstadter-io/geb/engine/utils"
)

type FileGenData struct {
	Dsl      string
	Gen      string
	File     string
	Data     interface{}
	Template *raymond.Template
	Outfile  string

	RepeatedContext interface{}
}

func (P *Project) Plan() error {
	logger.Info("Planning Project")

	P.register_partials()

	plans := []FileGenData{}

	// Loop over DSLs in the plans
	for d_key, D := range P.DslMap {
		logger.Info("    dsl: "+D.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Info("      gen: "+g_key, "gen_cfg", G.Config)

			// Render the normal templates
			for t_key, T := range G.Templates {
				t_ray := (*raymond.Template)(T)

				// build up the plan data struct
				fgd := FileGenData{
					Dsl:      d_key,
					Gen:      g_key,
					File:     t_key,
					Template: t_ray,
					Data:     P.Design,
					Outfile:  filepath.Join(P.Config.OutputDir, d_key, g_key, t_key),
				}
				logger.Info("        template file: "+t_key, "fgd", fgd)

				// add the plan to a linear list to be rendered
				plans = append(plans, fgd)
			} // End of normal template processing

			// only do repeats for actual dsls
			if D.Type != "dsl" {
				logger.Debug("       skipping dsl repeat: "+D.Type, "name", D.Name)
				continue
			}
			logger.Info("      doing dsl repeat: "+D.Type, "name", D.Name, "d_key", d_key)

			// Render the repeated templates
			// Get the root of the data to index into
			for k, _ := range P.Design.Dsl {
				logger.Debug("       - dsl keys", "key", k)
			}
			dsl_design, ok := P.Design.Dsl[d_key]
			if !ok {
				logger.Error("Did not find DSL data", "d_key", d_key)
				return errors.New("Unknown dsl design: " + d_key)
			}

			repeats := G.Config.Repeated
			logger.Info("Repeated found in config:", "count", len(repeats), "repeats", repeats)
			for _, R := range repeats {
				logger.Info("Processing Repeated Field " + R.Name)

				collection, err := utils.GetByPath(R.Field, (map[interface{}]interface{})(dsl_design))
				// collection, err := dsl_design.GetByPath(R.Field)
				if err != nil {
					return err
				}

				c_slice, ok := collection.([]interface{})
				if !ok {
					return errors.New("Collection is not a list: " + R.Field)
				}

				logger.Info("   Collection count", "collection", R.Field, "count", len(c_slice))
				for _, t_pair := range R.Templates {
					logger.Info("Looking for repeat template: ", "t_pair", t_pair, "in", G.Repeated)

					t_key := t_pair.In

					T, ok := G.Repeated[t_key]
					if !ok {
						return errors.New("Unknown repeat template: " + t_key)
					}
					t_ray := (*raymond.Template)(T)

					for idx, val := range c_slice {
						logger.Info("   Collection templates", "val", val, "count", len(R.Templates))

						local_ctx := val

						of_tpl_source := t_pair.Out
						tpl, err := raymond.Parse(of_tpl_source)
						if err != nil {
							return err
						}

						templates.AddHelpers(tpl)

						/*
							tpl_data := map[string]interface{}{
								"design": P.Design,
								"dsl":    dsl_design,
								"repeat": val,
							}
							OF_name, err := tpl.Exec(tpl_data)
						*/
						OF_name, err := tpl.Exec(val)
						if err != nil {
							return err
						}

						OF_name = strings.ToLower(OF_name)
						logger.Info("OFNAME", "name", OF_name)
						outfile := filepath.Join(P.Config.OutputDir, d_key, g_key, OF_name)

						// m_val := val.(map[interface{}]interface{})
						// elem_name := m_val["name"]
						// r_dir, r_file := filepath.Split(t_key)
						// of_name := fmt.Sprintf("%s-%s", elem_name, r_file)
						// outfile := filepath.Join(P.Config.OutputDir, d_key, g_key, r_dir, of_name)

						// build up the plan data struct
						fgd := FileGenData{
							Dsl:      d_key,
							Gen:      g_key,
							File:     t_key,
							Template: t_ray,
							Data:     P.Design,
							Outfile:  outfile,

							RepeatedContext: local_ctx,
						}
						logger.Info("        repeat file: "+t_key, "fgd", fgd, "index", idx)

						// add the plan to a linear list to be rendered
						plans = append(plans, fgd)

					}
				}

			}
			// End of repeated template processing

		} // End Generator loop

	} // End DSL loop

	P.Plans = plans

	return nil
}

func (P *Project) register_partials() {
	logger.Debug("Registering partials with templates and repeats")
	for d_key, D := range P.DslMap {
		logger.Debug("    dsl: "+D.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Debug("      gen: "+g_key, "gen_cfg", G.Config)

			// Register with the normal templates
			for _, T := range G.Templates {
				t_ray := (*raymond.Template)(T)
				for p_key, partial := range G.Partials {
					p_ray := (*raymond.Template)(partial)
					t_ray.RegisterPartialTemplate(p_key, p_ray)
				}
			}

			for _, R := range G.Repeated {
				t_ray := (*raymond.Template)(R)
				for p_key, partial := range G.Partials {
					p_ray := (*raymond.Template)(partial)
					t_ray.RegisterPartialTemplate(p_key, p_ray)
				}
			}

		}
	}

}
