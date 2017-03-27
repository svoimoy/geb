package unify
// package 

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.ibm.com/hofstadter-io/dotpath"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func Unify(parent string,path string,parentPath string,designData map[string]interface{}) (err error) {
	// HOFSTADTER_START Unify
	err = unify(parent, path, parentPath, designData)
	return err
	// HOFSTADTER_END   Unify
	return
}

/*
Where's your docs doc?!
*/
func unify(parent string,path string,parentPath string,designData interface{}) (err error) {
	// HOFSTADTER_START unify
	logger.Info("unify", "parent", parent, "path", path, "data", designData)
	path_flds := strings.Split(path, ".")
	path_len := len(path_flds)

	// Try to retrieve a name from the current object
	iname, err := dotpath.Get("name", designData, true)
	if err != nil {
		if !strings.Contains(err.Error(), "could not find 'name' in object") {
			return nil
			// return errors.Wrap(err, "in unify: "+path)
		}
	}

	name := ""

	if iname != nil {
		logger.Debug("Found a Name", "name", iname)
		// If we found a name, we found an object
		tname, ok := iname.(string)
		if !ok {
			return errors.New("in unify, obj '" + path + "' name is not a string")
		}
		name = tname

		pkg_path := strings.Join(path_flds[1:path_len-1], ".")
		pkg_path = strings.Replace(pkg_path, ".", "/", -1)

		switch vmap := designData.(type) {
		case map[string]interface{}:
			vmap["parent"] = parent
			vmap["parent_path"] = parentPath
			vmap["ctx_path"] = path
			vmap["pkg_path"] = pkg_path

		case map[interface{}]interface{}:
			vmap["parent"] = parent
			vmap["parent_path"] = parentPath
			vmap["ctx_path"] = path
			vmap["pkg_path"] = pkg_path

		default:
			return errors.New("in unify, named data is not a map")
		}

	}

	logger.Debug("Now inspecting obj", "data", designData)

	r_parent := parent
	r_parent_path := parentPath
	if name != "" {
		r_parent_path = path
		if parent != "" {
			r_parent = strings.Join([]string{parent, name}, ".")
		} else {
			r_parent = name
		}
	}

	logger.Debug("PARENT", "parent", parent, "r_parent", r_parent, "name", name)
	// now recurse
	switch D := designData.(type) {
	case map[string]interface{}:
		for key, val := range D {
			logger.Debug("  - inspecting...", "key", key, "val", val)
			r_path := strings.Join([]string{path, key}, ".")

			switch V := val.(type) {
			case map[string]interface{}:
				logger.Debug("Recursing  mS", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Debug("returning " + key)
					return errors.Wrap(err, "in unify: "+key)
				}
			case map[interface{}]interface{}:
				logger.Debug("Recursing  mS", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Debug("returning " + key)
					return errors.Wrap(err, "in unify: "+key)
				}
			case []interface{}:
				for idx, elem := range V {
					sidx := fmt.Sprint(idx)
					i_path := strings.Join([]string{r_path, sidx}, ".")
					logger.Debug("Recursing  mI []", "r_parent", r_parent, "r_path", i_path)
					err := unify(r_parent, i_path, r_parent_path, elem)
					if err != nil {
						logger.Debug("returning " + sidx)
						return errors.Wrap(err, "in unify: "+sidx)
					}
				}
			}
			logger.Debug("  - done inspecting...", "key", key, "val", val)
		}

	case map[interface{}]interface{}:
		for key, val := range D {
			logger.Debug("  - inspecting...", "key", key, "val", val)
			skey := fmt.Sprint(key)
			r_path := strings.Join([]string{path, skey}, ".")

			switch V := val.(type) {
			case map[string]interface{}:
				logger.Debug("Recursing  mI", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Debug("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			case map[interface{}]interface{}:
				logger.Debug("Recursing  mI", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Debug("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			case []interface{}:
				for idx, elem := range V {
					sidx := fmt.Sprint(idx)
					i_path := strings.Join([]string{r_path, sidx}, ".")
					logger.Debug("Recursing  mI []", "r_parent", r_parent, "r_path", i_path)
					err := unify(r_parent, i_path, r_parent_path, elem)
					if err != nil {
						logger.Debug("returning " + sidx)
						return errors.Wrap(err, "in unify: "+sidx)
					}
				}
			}
			logger.Debug("  - done inspecting...", "key", key, "val", val)
		}

	case []interface{}:
		for key, val := range D {
			logger.Debug("  - inspecting...", "key", key, "val", val)
			skey := fmt.Sprint(key)
			r_path := strings.Join([]string{path, skey}, ".")

			switch val.(type) {
			case map[string]interface{}:
				logger.Debug("Recursing  []i", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Debug("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			case map[interface{}]interface{}:
				logger.Debug("Recursing  []i", "r_parent", r_parent, "r_path", r_path)
				err := unify(r_parent, r_path, r_parent_path, val)
				if err != nil {
					logger.Debug("returning " + skey)
					return errors.Wrap(err, "in unify: "+skey)
				}
			}
			logger.Debug("  - done inspecting...", "key", key, "val", val)
		}

	default:
		logger.Debug("Not Recursing", "path", path)

		//	return errors.New("in unify, data is not a map")
	}

	logger.Info(" -- unify", "path", path)
	return nil
	// HOFSTADTER_END   unify
	return
}


// HOFSTADTER_BELOW
