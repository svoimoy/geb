package utils

import (
	"github.com/pkg/errors"
)

func Merge(orig, update interface{}) (interface{}, error) {

	if orig == nil {
		return update, nil
	}

	if update == nil {
		return orig, nil
	}

	// call the recursive merge
	return merge(orig, update)
}

func merge(orig, update interface{}) (interface{}, error) {

	logger.Warn("Merging", "orig", orig, "update", update)

	switch O := orig.(type) {

	case map[string]interface{}:
		U, ok := update.(map[string]interface{})
		if !ok {
			return nil, errors.New("update is not mS like orig")
		}
		logger.Warn("mS entering")
		for key, val := range U {
			logger.Error("in merge mS-U", "key", key, "val", val, "curr", O[key])
			if curr, exists := O[key]; exists {
				tmp, err := merge(curr, val)
				logger.Error("after merge mS", "tmp", tmp, "err", err)
				if err != nil {
					return nil, errors.Wrap(err, "in merge mS")
				}
				val = tmp
			}
			O[key] = val
		}
		logger.Warn("mS returning", "O", O)
		return O, nil

	case map[interface{}]interface{}:
		U, ok := update.(map[interface{}]interface{})
		if !ok {
			return nil, errors.New("update is not mI like orig")
		}
		logger.Warn("mI entering")
		for key, val := range U {
			logger.Error("in merge mI-U", "key", key, "val", val, "curr", O[key])
			if curr, exists := O[key]; exists {
				tmp, err := merge(curr, val)
				logger.Error("after merge mI", "tmp", tmp, "err", err)
				if err != nil {
					return nil, errors.Wrap(err, "in merge mI")
				}
				val = tmp
			}
			O[key] = val
		}
		logger.Warn("mI returning", "O", O)
		return O, nil

	case []interface{}:
		U, ok := update.([]interface{})
		if !ok {
			return nil, errors.New("update is not aI like orig")
		}

		logger.Warn("aI entering")
		// turn into maps
		OM := map[string]interface{}{}
		for i, elem := range O {
			switch E := elem.(type) {

			case map[string]interface{}:
				name, ok := E["name"]
				if !ok {
					return nil, errors.New("orig array objects must have names to be merged")
				}
				OM[name.(string)] = E

			case map[interface{}]interface{}:
				name, ok := E["name"]
				if !ok {
					return nil, errors.New("orig array objects must have names to be merged")
				}
				OM[name.(string)] = E

			case string:
				OM[E] = E

			default:
				logger.Error("orig unknown elem type in aI", "i", i, "elem", elem)
				return nil, errors.New("orig unknown elem type in aI")
			}
		}
		UM := map[string]interface{}{}
		for i, elem := range U {
			switch E := elem.(type) {

			case map[string]interface{}:
				name, ok := E["name"]
				if !ok {
					return nil, errors.New("orig array objects must have names to be merged")
				}
				UM[name.(string)] = E

			case map[interface{}]interface{}:
				name, ok := E["name"]
				if !ok {
					return nil, errors.New("orig array objects must have names to be merged")
				}
				UM[name.(string)] = E

			case string:
				UM[E] = E

			default:
				logger.Error("orig unknown elem type in aI", "i", i, "elem", elem)
				return nil, errors.New("orig unknown elem type in aI")
			}
		}

		// merge
		logger.Warn("aI")
		for key, val := range UM {
			if curr, exists := OM[key]; exists {
				tmp, err := merge(curr, val)
				logger.Error("in merge MS", "key", key, "val", val, "curr", curr)
				if err != nil {
					return nil, errors.Wrap(err, "in merge MS")
				}
				val = tmp
			}
			OM[key] = val
		}

		// turn back into array
		OA := []interface{}{}
		for _, val := range OM {
			OA = append(OA, val)
		}

		logger.Warn("aI returning", "OA", OA)
		return OA, nil

	default:
		return nil, errors.New("unmergable original")

	}

	logger.Crit("Shouldn't get here (end of merge function)")
	return orig, nil
}
