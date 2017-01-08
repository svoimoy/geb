package unify

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.ibm.com/hofstadter-io/dotpath"
	// HOFSTADTER_END   import
)

// This function adds meta data to the design_data.
// it may also verify some data or fields.
// Fields added are:
//
//  - parent
//  - ctx_path  (full design path to obj)
//  - pkg_path  (path to output objs relative to output)
//
//
func Unify(parent, path, parent_path string, design_data map[string]interface{}) error {
	logger.Crit("Unifying", "design_data", design_data)

	err := unify(parent, path, parent_path, design_data)
	logger.Crit("Post-unify", "design_data", design_data)

	return err
}

func unify(parent, path, parent_path string, design_data interface{}) error {
	logger.Error("unify", "parent", parent, "path", path, "data", design_data)
	path_flds := strings.Split(path, ".")
	path_len := len(path_flds)

	// Try to retrieve a name from the current object
	iname, err := dotpath.Get("name", design_data, true)
	if err != nil {
		if !strings.Contains(err.Error(), "could not find 'name' in object") {
			return errors.Wrap(err, "in unify: "+path)
		}
	}

	name := ""

	if iname != nil {
		logger.Warn("Found a Name", "name", iname)
		// If we found a name, we found an object
		tname, ok := iname.(string)
		if !ok {
			return errors.New("in unify, obj '" + path + "' name is not a string")
		}
		name = tname

		pkg_path := strings.Join(path_flds[1:path_len-1], ".")
		pkg_path = strings.Replace(pkg_path, ".", "/", -1)

		switch vmap := design_data.(type) {
		case map[string]interface{}:
			vmap["parent"] = parent
			vmap["parent_path"] = parent_path
			vmap["ctx_path"] = path
			vmap["pkg_path"] = pkg_path

		case map[interface{}]interface{}:
			vmap["parent"] = parent
			vmap["parent_path"] = parent_path
			vmap["ctx_path"] = path
			vmap["pkg_path"] = pkg_path

		default:
			return errors.New("in unify, named data is not a map")
		}

	}

	logger.Warn("Now inspecting obj", "data", design_data)

	r_parent := parent
	r_parent_path := parent_path
	if name != "" {
		r_parent_path = path
		if parent != "" {
			r_parent = strings.Join([]string{parent, name}, ".")
		} else {
			r_parent = name
		}
	}

	logger.Crit("PARENT", "parent", parent, "r_parent", r_parent, "name", name)
	// now recurse
	switch D := design_data.(type) {
	case map[string]interface{}:
		for key, val := range D {
			logger.Warn("  - inspecting...", "key", key, "val", val)
			r_path := strings.Join([]string{path, key}, ".")

			switch V := val.(type) {
			case map[string]interface{}:
				logger.Warn("Recursing  mS", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Crit("returning " + key)
					return errors.Wrap(err, "in unify: "+key)
				}
			case map[interface{}]interface{}:
				logger.Warn("Recursing  mS", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Crit("returning " + key)
					return errors.Wrap(err, "in unify: "+key)
				}
			case []interface{}:
				for idx, elem := range V {
					sidx := fmt.Sprint(idx)
					i_path := strings.Join([]string{r_path, sidx}, ".")
					logger.Warn("Recursing  mI []", "r_parent", r_parent, "r_path", i_path)
					err := unify(r_parent, i_path, r_parent_path, elem)
					if err != nil {
						logger.Crit("returning " + sidx)
						return errors.Wrap(err, "in unify: "+sidx)
					}
				}
			}
		}

	case map[interface{}]interface{}:
		for key, val := range D {
			logger.Warn("  - inspecting...", "key", key, "val", val)
			skey := fmt.Sprint(key)
			r_path := strings.Join([]string{path, skey}, ".")

			switch V := val.(type) {
			case map[string]interface{}:
				logger.Warn("Recursing  mI", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Crit("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			case map[interface{}]interface{}:
				logger.Warn("Recursing  mI", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Crit("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			case []interface{}:
				for idx, elem := range V {
					sidx := fmt.Sprint(idx)
					i_path := strings.Join([]string{r_path, sidx}, ".")
					logger.Warn("Recursing  mI []", "r_parent", r_parent, "r_path", i_path)
					err := unify(r_parent, i_path, r_parent_path, elem)
					if err != nil {
						logger.Crit("returning " + sidx)
						return errors.Wrap(err, "in unify: "+sidx)
					}
				}
			}
		}

	case []interface{}:
		for key, val := range D {
			logger.Warn("  - inspecting...", "key", key, "val", val)
			skey := fmt.Sprint(key)
			r_path := strings.Join([]string{path, skey}, ".")

			switch val.(type) {
			case map[string]interface{}:
				logger.Warn("Recursing  []i", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Crit("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			case map[interface{}]interface{}:
				logger.Warn("Recursing  []i", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Crit("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			}
		}

	default:
		logger.Warn("Not Recursing", "path", path)

		//	return errors.New("in unify, data is not a map")
	}

	logger.Error(" -- unify")
	return nil
}
