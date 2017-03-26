package design

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"strings"

	"github.ibm.com/hofstadter-io/dotpath"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import
)

func (d *Design) store_dsl_design(relative_path, dsl, name string, design interface{}) error {
	logger.Info("    - storing dsl", "dsl", dsl, "name", name, "rel_path", relative_path)
	logger.Debug("          with", "design", design)

	fields := strings.Split(relative_path, "/")
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
	curr_path := []string{}
	for i, F := range fields {
		foundSameDsl := false

		sameDslPath := append(curr_path, dsl, name)
		dsl_path := strings.Join(sameDslPath, ".")

		no_solo_array := true
		obj, err := dotpath.Get(dsl_path, d.Dsl, no_solo_array)
		if err != nil {
			logger.Debug("in store_dsl_design", "error", err)
			// return errors.Wrap(err, "Error during path search in store_dsl_design.")
		}
		if obj != nil {
			// we found a same dsl
			foundSameDsl = true
			F0 = strings.Split(dsl_path, ".")[0]
		}

		logger.Info("store_dsl_design", "curr_path", curr_path, "dsl_path", dsl_path, "i", i, "rel_path", relative_path, "dsl", dsl, "name", name, "match", foundSameDsl)

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

	if _, ok := d.Dsl[F0]; !ok {
		d.Dsl[F0] = insert[F0]
		logger.Info("new dsl stored", "d.Dsl", d.Dsl)
	} else {
		logger.Info("merge...", "d.Dsl", d.Dsl, "update", insert)

		merged, merr := utils.Merge(d.Dsl, insert)
		if merr != nil {
			return errors.Wrap(merr, "in store_dsl_design")
		}
		logger.Info("result...", "merged", merged)

		d.Dsl[F0] = merged.(map[string]interface{})[F0]
		logger.Debug("final merge", "d.Dsl", d.Dsl)
	}

	return nil
}

