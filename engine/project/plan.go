package project

import (
	"github.com/pkg/errors"
	"os"
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
	P.add_template_helpers()

	plans := []FileGenData{}

	logger.Info("Making the plans")

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

			// Start of repeat processing section:
			//
			// only do repeats for actual dsls and
			//   when there are repeats
			//
			repeats := G.Config.Repeated
			if D.Type != "dsl" || len(repeats) == 0 {
				logger.Debug("       skipping dsl repeat: "+D.Type, "name", D.Name)
				continue
			}
			logger.Info("Repeated found in config:", "count", len(repeats), "repeats", repeats)
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

			for _, R := range repeats {
				logger.Info("Processing Repeated Field: '" + R.Name + "'")

				// look up field
				collection, err := utils.GetByPath(R.Field, (map[interface{}]interface{})(dsl_design))
				if err != nil {
					return errors.Errorf("looking up by path:  repeat(" + R.Name + ")  path" + R.Field + ")")
				}

				c_slice, ok := collection.([]interface{})
				if !ok {
					return errors.New("Collection is not a list: " + R.Field)
				}

				logger.Info("   Collection count", "collection", R.Field, "count", len(c_slice))
				for _, t_pair := range R.Templates {
					// fmt.Println("AAAAAAAA = ", I)
					logger.Info("    Looking for repeat template: ", "t_pair", t_pair, "in", G.Repeated)

					t_key := t_pair.In

					T, ok := G.Repeated[t_key]
					if !ok {
						// fmt.Println("    XXXX = ", t_key)
						return errors.New("Unknown repeat template: " + t_key)
					}
					t_ray := (*raymond.Template)(T)
					logger.Debug("        found repeat template: ", "repeat", R.Name, "in", t_key)
					// fmt.Println("BBBBBBBB = ", I)

					os.Stdout.Sync()

					for idx, val := range c_slice {

						// fmt.Println("    AAAA = ", idx, val)
						local_ctx := val
						logger.Debug("     context", "val", local_ctx, "idx", idx)
						os.Stdout.Sync()

						of_tpl_source := t_pair.Out
						tpl, err := raymond.Parse(of_tpl_source)
						if err != nil {
							return err
						}
						// fmt.Println("    BBBB = ", idx, val)

						templates.AddHelpers(tpl)
						OF_name, err := tpl.Exec(val)
						if err != nil {
							return err
						}
						// fmt.Println("    TTTT = ", idx, OF_name)

						OF_name = strings.ToLower(OF_name)
						logger.Info("OFNAME", "name", OF_name)
						outfile := filepath.Join(P.Config.OutputDir, d_key, g_key, OF_name)

						// fmt.Println("    CCCC = ", I)
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

	P.Plans = plans

	return nil
}
