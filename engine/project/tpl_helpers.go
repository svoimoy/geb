package project

import (
	"fmt"
	"strings"

	"github.com/aymerick/raymond"
	"github.ibm.com/hofstadter-io/dotpath"
)

func (P *Project) register_partials() {
	logger.Debug("Registering partials with templates and repeats")
	for d_key, D := range P.DslMap {
		logger.Debug("    dsl: "+D.Config.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Debug("      gen: "+g_key, "gen_cfg", G.Config)

			// Register partials with the normal templates
			for _, T := range G.Templates {
				t_ray := (*raymond.Template)(T)
				for p_key, partial := range G.Partials {
					p_ray := (*raymond.Template)(partial)
					t_ray.RegisterPartialTemplate(p_key, p_ray)
				}

				// This nasty for loop nesting is to add the global partials to the templates
				// Loop over each generator in the current DSL
				for d2_key, D2 := range P.DslMap {
					// Loop over each generator in the current DSL
					for g2_key, G2 := range D2.Generators {
						for p2_key, partial2 := range G2.Partials {
							p2_ray := (*raymond.Template)(partial2)
							p2_tkey := strings.Join([]string{d2_key, g2_key, p2_key}, "/")
							logger.Debug("adding global partial " + p2_tkey)
							t_ray.RegisterPartialTemplate(p2_tkey, p2_ray)
						}
					}
				}

			} // end loop over normal templates

			// Register partials with repeated templates
			for _, R := range G.Repeated {
				// the real template object
				t_ray := (*raymond.Template)(R)

				// register the local generator partials
				for p_key, partial := range G.Partials {
					p_ray := (*raymond.Template)(partial)
					t_ray.RegisterPartialTemplate(p_key, p_ray)
				}

				// register the global partials
				// This nasty for loop nesting is to add the global partials to the templates
				// Loop over each DSL in the current Project
				for d2_key, D2 := range P.DslMap {
					// Loop over each generator in the current DSL
					for g2_key, G2 := range D2.Generators {
						for p2_key, partial2 := range G2.Partials {
							p2_ray := (*raymond.Template)(partial2)
							p2_tkey := strings.Join([]string{d2_key, g2_key, p2_key}, "/")
							logger.Debug("adding global partial " + p2_tkey)
							t_ray.RegisterPartialTemplate(p2_tkey, p2_ray)
						}
					}
				}

			} // end loop over repeated templates

		}
	}

}

func (P *Project) add_template_helpers() {

	logger.Debug("Registering partials with templates and repeats")
	for d_key, D := range P.DslMap {
		logger.Debug("    dsl: "+D.Config.Name, "key", d_key)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Debug("      gen: "+g_key, "gen_cfg", G.Config)

			// Register with the normal templates
			for _, template := range G.Templates {
				ray := (*raymond.Template)(template)
				P.register_template_helpers(ray)
			}

			for _, repeated := range G.Repeated {
				ray := (*raymond.Template)(repeated)
				P.register_template_helpers(ray)
			}
			for _, partial := range G.Partials {
				ray := (*raymond.Template)(partial)
				P.register_template_helpers(ray)
			}

		}
	}

}

func (P *Project) register_template_helpers(tpl *raymond.Template) {
	tpl.RegisterHelper("dotpath", P.tpl_helper_dotpath)
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

// data optional argument defaults to DSL
func (P *Project) tpl_helper_dotpath_type(path string, no_solo_array bool, options *raymond.Options) interface{} {
	data := P.Design.Type

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
