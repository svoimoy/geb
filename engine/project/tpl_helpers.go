package project

import (
	"fmt"
	"github.com/aymerick/raymond"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
)

func (P *Project) register_partials() {
	logger.Debug("Registering partials with templates and repeats")
	for d_key, D := range P.DslMap {
		logger.Debug("    dsl: "+D.Config.Name, "key", d_key)

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
func (P *Project) tpl_helper_get_elem_by_name(path, name string, options *raymond.Options) interface{} {
	hash := options.Hash()
	data, ok := hash["data"]
	if !ok {
		data = P.Design.Dsl
	} else if data == nil {
		return options.FnWith("Nil data supplied" + path)
	}

	obj, err := utils.GetByPath(path, data)
	if err != nil {
		return options.FnWith("Error during path search: " + err.Error())
	}
	// obj := options.Eval(data, path)
	if obj == nil {
		return options.FnWith("Path not found: " + path + fmt.Sprintf("\n%+v", data))
	}

	return options.FnWith(obj)

	list, ok := obj.([]interface{})
	if !ok {
		return options.FnWith("Path is not a list: " + path)
	}

	for I, e := range list {
		switch E := e.(type) {
		case map[string]interface{}:
			n, ok := E["name"]
			if !ok {
				return options.FnWith(fmt.Sprintf("Path list element is missing name: %s %d", path, I))
			}
			if n == name {
				return options.FnWith(e)
			}

		case map[interface{}]interface{}:
			n, ok := E["name"]
			if !ok {
				return options.FnWith(fmt.Sprintf("Path list element is missing name: %s %d", path, I))
			}
			if n == name {
				return options.FnWith(e)
			}

		default:
			return options.FnWith("Path list elements are not objects: " + path)
		}
	}
	return options.FnWith("Did not find object with name '" + name + "' in path: " + path)
}
