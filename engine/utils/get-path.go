package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func GetByPath(path string, data interface{}) (interface{}, error) {
	paths := strings.Split(path, ".")
	if len(paths) < 1 {
		return nil, errors.New("Bad path supplied: " + path)
	}

	// fmt.Println("GETPATH:", path, paths, data)

	return get_by_path(0, paths, data)
}

func GetByPathSlice(path []string, data interface{}) (interface{}, error) {
	return get_by_path(0, path, data)
}

func get_by_path(IDX int, paths []string, data interface{}) (interface{}, error) {
	header := fmt.Sprintf("get_by_path:  %d  %v  in:\n%+v\n\n", IDX, paths, data)
	// fmt.Println(header)
	logger.Info(header)

	P := paths[IDX]
	path_str := strings.Join(paths[:IDX+1], ".")

	switch T := data.(type) {

	case map[string]interface{}:
		val, ok := T[P]
		if !ok {
			return nil, errors.New("could not find '" + P + "' in object")
		}
		add_parent_and_path(val, T, path_str)
		if len(paths) == IDX+1 {
			return val, nil
		}
		ret, err := get_by_path(IDX+1, paths, val)
		if err != nil {
			return nil, errors.Wrapf(err, "from object "+P)
		}
		return ret, nil

	case map[interface{}]interface{}:
		val, ok := T[P]
		if !ok {
			return nil, errors.New("could not find '" + P + "' in object")
		}
		add_parent_and_path(val, T, path_str)
		if len(paths) == IDX+1 {
			return val, nil
		}
		ret, err := get_by_path(IDX+1, paths, val)
		if err != nil {
			return nil, errors.Wrapf(err, "from object "+P)
		}
		return ret, nil

	case []interface{}:
		logger.Info("Processing Slice", "paths", paths, "T", T)
		subs := []interface{}{}
		if len(paths) == IDX+1 {
			for _, elem := range T {
				logger.Info("    - elem", "elem", elem, "paths", paths, "P", P, "elem", elem)
				switch V := elem.(type) {

				case map[string]interface{}:
					logger.Debug("        map[string]")
					val, ok := V[P]
					if !ok {
						logger.Debug("could not find '" + P + "' in object")
						continue
					}

					// accumulate based on type (slice or not)
					switch a_val := val.(type) {

					case []interface{}:
						logger.Debug("Adding vals", "val", a_val)
						subs = append(subs, a_val...)

					default:
						logger.Debug("Adding val", "val", a_val)
						subs = append(subs, a_val)
					}

				case map[interface{}]interface{}:
					logger.Debug("        map[iface]", "P", P, "V", V, "paths", paths)
					val, ok := V["name"]
					if !ok {
						logger.Debug("could not find '" + P + "' in object")
						continue
					}

					if val == P {

						// accumulate based on type (slice or not)
						switch a_val := val.(type) {

						case []interface{}:
							logger.Debug("Adding vals", "val", a_val)
							subs = append(subs, a_val...)

						default:
							logger.Debug("Adding val", "val", a_val)
							subs = append(subs, a_val)
						}
						return V, nil
					}

				default:
					str := fmt.Sprintf("%+v", reflect.TypeOf(V))
					return nil, errors.New("element not an object type: " + str)

				}
			}
		} else {
			for _, elem := range T {
				val, err := get_by_path(IDX, paths, elem)
				if err != nil {
					// in this case, only some of the sub.paths.elements may be found
					// this err path should be expanded to check for geb error types
					logger.Debug(err.Error())
					continue
				}
				switch V := val.(type) {

				case []interface{}:
					logger.Debug("Adding vals", "val", V)
					subs = append(subs, V...)

				default:
					logger.Debug("Adding val", "val", V)
					subs = append(subs, V)

				}
			}
		}

		return subs, nil

	default:
		str := fmt.Sprintf("%+v", reflect.TypeOf(data))
		return nil, errors.New("unknown data object type: " + str)

	} // END of type switch

}

func add_parent_and_path(child interface{}, parent interface{}, path string) (interface{}, error) {
	logger.Info("adding parent to child", "child", child, "parent", parent, "path", path)
	parent_ref := "unknown-parent"
	switch P := parent.(type) {

	case map[string]interface{}:
		p_ref, ok := P["name"]
		if !ok {
			return nil, errors.Errorf("parent does not have name: %+v", parent)
		}
		parent_ref = p_ref.(string)
	case map[interface{}]interface{}:
		p_ref, ok := P["name"]
		if !ok {
			return nil, errors.Errorf("parent does not have name: %+v", parent)
		}
		parent_ref = p_ref.(string)

	default:
		str := fmt.Sprintf("%+v", reflect.TypeOf(parent))
		return nil, errors.New("unknown parent object type: " + str)

	}

	switch C := child.(type) {

	case map[string]interface{}:
		C["parent"] = parent_ref
		C["path"] = path
	case map[interface{}]interface{}:
		C["parent"] = parent_ref
		C["path"] = path

	case []interface{}:
		for _, elem := range C {
			switch E := elem.(type) {
			case map[string]interface{}:
				E["parent"] = parent_ref
				E["path"] = path
			case map[interface{}]interface{}:
				E["parent"] = parent_ref
				E["path"] = path
			default:
				str := fmt.Sprintf("in slice of %+v", reflect.TypeOf(E))
				return nil, errors.New("element not an object type: " + str)
			}
		}

	default:
		str := fmt.Sprintf("%+v", reflect.TypeOf(C))
		return nil, errors.New("unknown data object type: " + str)

	}
	return child, nil
}
