package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func GetByPath(path string, data interface{}) (interface{}, error) {
	paths := strings.Split(path, ".")
	if len(paths) < 1 {
		return nil, errors.New("Bad path supplied: " + path)
	}

	return get_by_path(paths, data)
}

func GetByPathSlice(path []string, data interface{}) (interface{}, error) {
	return get_by_path(path, data)
}

func get_by_path(paths []string, data interface{}) (interface{}, error) {
	logger.Info(fmt.Sprintf("get_by_path:  %v  in:\n%+v\n\n", paths, data))

	P := paths[0]

	switch T := data.(type) {

	case map[string]interface{}:
		val, ok := T[P]
		if !ok {
			return nil, errors.New("could not find '" + P + "' in object")
		}
		if len(paths) == 1 {
			return val, nil
		}
		return get_by_path(paths[1:], val)

	case map[interface{}]interface{}:
		val, ok := T[P]
		if !ok {
			return nil, errors.New("could not find '" + P + "' in object")
		}
		if len(paths) == 1 {
			return val, nil
		}
		return get_by_path(paths[1:], val)

	case []interface{}:
		logger.Info("Processing Slice", "paths", paths, "T", T)
		subs := []interface{}{}
		if len(paths) == 1 {
			for _, elem := range T {
				logger.Info("    - elem", "elem", elem)
				switch V := elem.(type) {

				case map[string]interface{}:
					val, ok := V[P]
					if !ok {
						logger.Debug("could not find '" + P + "' in object")
						continue
					}
					switch a_val := val.(type) {
					case []interface{}:
						logger.Debug("Adding vals", "val", a_val)
						subs = append(subs, a_val...)
					default:
						logger.Debug("Adding val", "val", a_val)
						subs = append(subs, a_val)
					}

				case map[interface{}]interface{}:
					val, ok := V[P]
					if !ok {
						logger.Debug("could not find '" + P + "' in object")
						continue
					}
					switch a_val := val.(type) {
					case []interface{}:
						logger.Debug("Adding vals", "val", a_val)
						subs = append(subs, a_val...)
					default:
						logger.Debug("Adding val", "val", a_val)
						subs = append(subs, a_val)
					}

				default:
					str := fmt.Sprintf("%+v", reflect.TypeOf(V))
					return nil, errors.New("unknown data object type: " + str)

				}
			}
		} else {
			for _, elem := range T {
				val, err := get_by_path(paths, elem)
				if err != nil {
					logger.Debug(err.Error())
					continue
					return nil, err
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
