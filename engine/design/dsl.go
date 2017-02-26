package design

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"strings"

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
	for _, F := range fields {
		if F != "" {
			tmp := make(map[string]interface{})
			dd_map[F] = tmp
			dd_map = tmp
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
