package design

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"strings"

	"github.com/hofstadter-io/dotpath"

	// "github.com/hofstadter-io/geb/engine/utils"
	"github.com/hofstadter-io/data-utils/manip"

	"fmt"
	"github.com/mohae/deepcopy"
	"github.com/go-test/deep"

	// HOFSTADTER_END   import
)

/*
Name:      store
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
func (D *Design) storeDslDesign(relativePath string, dsl string, name string, design interface{}) (err error) {
	// HOFSTADTER_START storeDslDesign
	logger.Info("    - storing dsl", "dsl", dsl, "name", name, "rel_path", relativePath)
	logger.Debug("          with", "design", design)

	fields := strings.Split(relativePath, "/")
	F0 := fields[0]
	logger.Debug("Fields", "fields", fields)
	if F0 == "" {
		F0 = dsl
	}
	logger.Debug("F0: '" + F0 + "'")

	// This block builds up the object to insert
	// - from the outermost map
	// - through the namespace fields
	// - and finally the design itself
	insert := make(map[string]interface{})
	dd_map := insert
	for i, F := range fields {
		foundSameDsl := false
		curr_path := make([]string, len(fields[:i]), len(fields))
		copy(curr_path, fields[:i])

		sameDslPath := append(curr_path, dsl, name)
		// sameDslPath = append(sameDslPath, name)
		dsl_path := strings.Join(sameDslPath, ".")

		no_solo_array := true
		obj, err := dotpath.Get(dsl_path, D.Dsl, no_solo_array)
		if err != nil {
			logger.Debug("in store_dsl_design", "error", err, "dsl_path", dsl_path, "curr_path", curr_path)
			// return errors.Wrap(err, "Error during path search in store_dsl_design.")
		}
		if obj != nil {
			// we found a same dsl
			logger.Info("Found same dsl", "dsl", dsl)
			foundSameDsl = true
			F0 = strings.Split(dsl_path, ".")[0]
		}

		logger.Info("store_dsl_design", "F", F, "curr_path", curr_path, "dsl_path", dsl_path, "i", i, "rel_path", relativePath, "dsl", dsl, "name", name, "match", foundSameDsl)

		cmap := make(map[string]interface{})
		if foundSameDsl {
			// short circuit insertion path
			break
		}

		// add the current path part to the insert path
		if F != "" {
			dd_map[F] = cmap
			dd_map = cmap
		}
	}
	dd_map[dsl] = design
	logger.Info("Design", "name", name, "design", design, "map", dd_map, "insert", insert)

	if _, ok := D.Dsl[F0]; !ok {
		D.Dsl[F0] = insert[F0]
		logger.Info("new dsl stored", "D.Dsl", D.Dsl)
	} else {
		logger.Info("merge...", "D.Dsl", D.Dsl, "update", insert)

		merged, merr := manip.Merge(D.Dsl, insert)
		if merr != nil {
			return errors.Wrap(merr, "in storeDslDesign")
		}
		logger.Info("result...", "merged", merged)

		D.Dsl[F0] = merged.(map[string]interface{})[F0]
		logger.Debug("final merge", "D.Dsl", D.Dsl)
	}

	return nil
	// HOFSTADTER_END   storeDslDesign
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) storeDataDesign(relativePath string, name string, design interface{}) (err error) {
	// HOFSTADTER_START storeDataDesign
	// HOFSTADTER_END   storeDataDesign
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) storeTypeDesign(relativePath string, name string, design interface{}) (err error) {
	// HOFSTADTER_START storeTypeDesign
	logger.Info("    - storing type", "name", name, "rel_path", relativePath)
	logger.Debug("          with", "design", design)

	fields := strings.Split(relativePath, "/")
	F0 := fields[0]
	logger.Debug("Fields", "fields", fields)
	if F0 == "" {
		F0 = name
	}

	// This block builds up the object to insert
	// - from the outermost map
	// - through the namespace fields
	// - and finally the design itself
	insert := make(map[string]interface{})
	dd_map := insert
	for _, F := range fields {
		if F != "" {
			tmp := make(map[string]interface{})
			dd_map[F] = tmp
			dd_map = tmp
		}
	}
	dd_map[name] = design
	logger.Info("Design", "name", name, "design", design, "map", dd_map, "insert", insert)

	if _, ok := D.Type[F0]; !ok {
		D.Type[F0] = insert[F0]
		logger.Info("new type stored", "D.Type", D.Type)
	} else {
		logger.Info("merge...", "D.Type", D.Type, "update", insert)

		orig := deepcopy.Copy(D.Type)

		merged, merr := manip.Merge(D.Type, insert)
		if merr != nil {
			return errors.Wrap(merr, "in storeTypeDesign")
		}

		equal := deep.Equal(orig, D.Type)
		fmt.Println("DESIGN == Post-Subdesign: ", len(equal))

		logger.Info("result...", "merged", merged)
		D.Type[F0] = merged.(map[string]interface{})[F0]
		logger.Debug("final merge", "D.Type", D.Type)
	}
	logger.Debug("       - " + F0)

	return nil

	// HOFSTADTER_END   storeTypeDesign
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) storePackageDesign(relativePath string, name string, design interface{}) (err error) {
	// HOFSTADTER_START storePackageDesign
	logger.Info("    - storing pkg", "name", name, "rel_path", relativePath)
	logger.Debug("          with", "design", design)

	fields := strings.Split(relativePath, "/")
	F0 := fields[0]
	logger.Debug("Fields", "fields", fields)
	if F0 == "" {
		F0 = name
	}

	// This block builds up the object to insert
	// - from the outermost map
	// - through the namespace fields
	// - and finally the design itself
	insert := make(map[string]interface{})
	dd_map := insert
	for _, F := range fields {
		if F != "" {
			tmp := make(map[string]interface{})
			dd_map[F] = tmp
			dd_map = tmp
		}
	}
	dd_map[name] = design
	logger.Info("Design", "name", name, "design", design, "map", dd_map, "insert", insert)

	if _, ok := D.Pkg[F0]; !ok {
		D.Pkg[F0] = insert[F0]
		logger.Debug("new pkg stored", "D.Pkg", D.Pkg)
	} else {
		logger.Info("merge...", "D.Pkg", D.Pkg, "update", insert)

		merged, merr := manip.Merge(D.Pkg, insert)
		if merr != nil {
			logger.Warn("Error merging", "error", merr)
			return errors.Wrap(merr, "in storePakcageDesign")
		}
		logger.Info("result...", "merged", merged)

		D.Pkg[F0] = merged.(map[string]interface{})[F0]
		logger.Debug("final merge", "D.Pkg", D.Pkg)
	}

	return nil
	// HOFSTADTER_END   storePackageDesign
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) storeCustomDesign(relativePath string, name string, design interface{}) (err error) {
	// HOFSTADTER_START storeCustomDesign
	logger.Info("    - storing custom", "name", name, "rel_path", relativePath)
	logger.Debug("          with", "design", design)

	fields := strings.Split(relativePath, "/")
	F0 := fields[0]
	logger.Debug("Fields", "fields", fields)
	if F0 == "" {
		F0 = name
	}

	// This block builds up the object to insert
	// - from the outermost map
	// - through the namespace fields
	// - and finally the design itself
	insert := make(map[string]interface{})
	dd_map := insert
	for _, F := range fields {
		if F != "" {
			tmp := make(map[string]interface{})
			dd_map[F] = tmp
			dd_map = tmp
		}
	}
	dd_map[name] = design
	logger.Info("Design", "name", name, "design", design, "map", dd_map, "insert", insert)

	if _, ok := D.Custom[F0]; !ok {
		D.Custom[F0] = insert[F0]
		logger.Debug("new custom data stored", "D.Custom", D.Custom)
	} else {
		logger.Info("merge...", "D.Custom", D.Custom, "update", insert)

		merged, merr := manip.Merge(D.Custom, insert)
		if merr != nil {
			return errors.Wrap(merr, "in storeCustomDesign")
		}
		logger.Info("result...", "merged", merged)

		D.Custom[F0] = merged.(map[string]interface{})[F0]
		logger.Debug("final merge", "D.Custom", D.Custom)
	}

	return nil
	// HOFSTADTER_END   storeCustomDesign
	return
}

// HOFSTADTER_BELOW
