package project

import (
	// HOFSTADTER_START import
	"fmt"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/hofstadter-io/dotpath"
	// HOFSTADTER_END   import
)

/*
Name:      template-helpers
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
func (P *Project) registerPartials() {
	// HOFSTADTER_START registerPartials
	logger.Debug("Registering partials with templates and repeats")
	for d_key, D := range P.DslMap {
		logger.Debug("    dsl: "+D.Config.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Debug("      gen: "+g_key, "gen_cfg", G.Config)

			// Register partials with repeated templates
			for _, R := range G.Templates {
				// the real template object
				t_ray := (*raymond.Template)(R.Template)

				// register the local generator partials
				for p_key, partial := range G.Partials {
					p_ray := (*raymond.Template)(partial.Template)
					t_ray.RegisterPartialTemplate(p_key, p_ray)
				}

				// register the global partials
				// This nasty for loop nesting is to add the global partials to the templates
				// Loop over each DSL in the current Project
				for d2_key, D2 := range P.DslMap {
					// Loop over each generator in the current DSL
					for g2_key, G2 := range D2.Generators {
						for p2_key, partial2 := range G2.Partials {
							p2_ray := (*raymond.Template)(partial2.Template)
							p2_tkey := strings.Join([]string{d2_key, g2_key, p2_key}, "/")
							logger.Debug("adding global partial " + p2_tkey)
							t_ray.RegisterPartialTemplate(p2_tkey, p2_ray)
						}
					}
				}

			} // end loop over repeated templates

		}
	}

	// HOFSTADTER_END   registerPartials
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) addTemplateHelpers() {
	// HOFSTADTER_START addTemplateHelpers
	logger.Debug("Registering partials with templates and repeats")
	for d_key, D := range P.DslMap {
		logger.Debug("    dsl: "+D.Config.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Debug("      gen: "+g_key, "gen_cfg", G.Config)

			// Register partials with the templates
			for _, template := range G.Templates {
				ray := (*raymond.Template)(template.Template)
				P.registerTemplateHelpers(ray)
			}

			// Register partials with the partials
			for _, partial := range G.Partials {
				ray := (*raymond.Template)(partial.Template)
				P.registerTemplateHelpers(ray)
			}

			// Register helpers with the designs
			for _, design := range G.Designs {
				ray := (*raymond.Template)(design.Template)
				P.registerTemplateHelpers(ray)
			}
		}
	}
	// HOFSTADTER_END   addTemplateHelpers
	return
}

// HOFSTADTER_BELOW

func (P *Project) registerTemplateHelpers(tpl *raymond.Template) {
	tpl.RegisterHelper("dotpath", P.tpl_helper_dotpath)
	tpl.RegisterHelper("getdesign", P.tpl_helper_dotpath_design)
	tpl.RegisterHelper("gettype", P.tpl_helper_dotpath_type)
	tpl.RegisterHelper("getdsl", P.tpl_helper_dotpath_dsl)

	tpl.RegisterHelper("get_obj_by_path", P.tpl_helper_get_obj_by_path)
	tpl.RegisterHelper("get_elem_by_name", P.tpl_helper_get_elem_by_name)
}

// data optional argument defaults to TYPE
func (P *Project) tpl_helper_get_obj_by_path(path string, options *raymond.Options) interface{} {
	hash := options.Hash()
	data, ok := hash["data"]
	if !ok {
		data = P.Design.Type
	} else if data == nil {
		return options.FnWith("Path not found: " + path)
	}

	// obj, err := GetByPath(path, data)
	// if err != nil {
	// 	return options.FnWith("Error during path search: " + err.Error())
	// }

	obj := options.Eval(data, path)
	if obj == nil {
		return options.FnWith("Path not found: " + path)
	}
	return options.FnWith(obj)
}

// data optional argument defaults to DSL
func (P *Project) tpl_helper_get_elem_by_name(path, name string, no_solo_array bool, options *raymond.Options) interface{} {
	hash := options.Hash()
	data, ok := hash["data"]
	if !ok {
		data = P.Design.Dsl
	} else if data == nil {
		return options.FnWith("Nil data supplied" + path)
	}

	obj, err := dotpath.Get(path, data, no_solo_array)
	if err != nil {
		return options.FnWith("Error during path search: " + err.Error())
	}
	// obj := options.Eval(data, path)
	if obj == nil {
		return options.FnWith("Path not found: " + path + fmt.Sprintf("\n%+v", data))
	}

	return options.FnWith(obj)
}

// data optional argument defaults to DSL
func (P *Project) tpl_helper_dotpath(path string, data interface{}, no_solo_array bool, options *raymond.Options) interface{} {
	if data == nil {
		return options.FnWith("Nil data supplied" + path)
	}

	obj, err := dotpath.Get(path, data, no_solo_array)
	if err != nil {
		return options.FnWith("Error during path search: " + err.Error())
	}
	// obj := options.Eval(data, path)
	if obj == nil {
		return options.FnWith("Path not found: " + path + fmt.Sprintf("\n%+v", data))
	}

	return options.FnWith(obj)
}

// data optional argument defaults to Design
func (P *Project) tpl_helper_dotpath_design(path string, no_solo_array bool, options *raymond.Options) interface{} {
	data := P.Design

	obj, err := data.GetByPath(path)
	if err != nil {
		return options.FnWith("Error during path search: " + err.Error())
	}
	if obj == nil {
		return options.FnWith("Path not found: " + path + fmt.Sprintf("\n%+v", data))
	}

	return options.FnWith(obj)
}

// data optional argument defaults to DSL
func (P *Project) tpl_helper_dotpath_type(path string, no_solo_array bool, options *raymond.Options) interface{} {
	data := P.Design.Type

	if path == "" {
		return options.FnWith("Empty path!!")
	}
	if strings.HasPrefix(path, "type.") {
		path = strings.TrimPrefix(path, "type.")
	}

	obj, err := dotpath.Get(path, data, no_solo_array)
	if err != nil {
		return options.FnWith("Error during path search: " + err.Error())
	}
	if obj == nil {
		return options.FnWith("Path not found: " + path + fmt.Sprintf("\n%+v", data))
	}

	return options.FnWith(obj)
}

// data optional argument defaults to DSL
func (P *Project) tpl_helper_dotpath_dsl(path string, no_solo_array bool, options *raymond.Options) interface{} {
	data := P.Design.Dsl

	obj, err := dotpath.Get(path, data, no_solo_array)
	if err != nil {
		return options.FnWith("Error during path search: " + err.Error())
	}
	if obj == nil {
		return options.FnWith("Path not found: " + path + fmt.Sprintf("\n%+v", data))
	}

	return options.FnWith(obj)
}
