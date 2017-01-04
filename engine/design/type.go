package design

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"strings"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import
)

func (d *Design) store_type_design(relative_path, name string, design interface{}) error {
	logger.Info("    - storing type", "name", name, "rel_path", relative_path)
	logger.Debug("          with", "design", design)

	fields := strings.Split(relative_path, "/")
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

	if _, ok := d.Type[F0]; !ok {
		d.Type[F0] = insert[F0]
		logger.Info("new type stored", "d.Type", d.Type)
	} else {
		logger.Info("merge...", "d.Type", d.Type, "update", insert)

		merged, merr := utils.Merge(d.Type, insert)
		if merr != nil {
			return errors.Wrap(merr, "in store_type_design")
		}
		logger.Info("result...", "merged", merged)
		d.Type[F0] = merged.(map[string]interface{})[F0]
		logger.Debug("final merge", "d.Type", d.Type)
	}
	logger.Debug("       - " + F0)

	return nil
}
