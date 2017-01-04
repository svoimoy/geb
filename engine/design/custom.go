package design

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"strings"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import
)

func (d *Design) store_custom_design(relative_path, name string, design interface{}) error {
	logger.Info("    - storing custom", "name", name, "rel_path", relative_path)
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

	if _, ok := d.Custom[F0]; !ok {
		d.Custom[F0] = insert[F0]
		logger.Debug("new custom data stored", "d.Custom", d.Custom)
	} else {
		logger.Info("merge...", "d.Custom", d.Custom, "update", insert)

		merged, merr := utils.Merge(d.Custom, insert)
		if merr != nil {
			return errors.Wrap(merr, "in store_pkg_design")
		}
		logger.Info("result...", "merged", merged)

		d.Custom[F0] = merged.(map[string]interface{})[F0]
		logger.Debug("final merge", "d.Custom", d.Custom)
	}

	return nil
}
