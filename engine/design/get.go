package design

import (
	"errors"
	"strings"

	"github.com/hofstadter-io/geb/engine/utils"
)

func (D *Design) Get(path string) (interface{}, error) {
	paths := strings.Split(path, ".")
	if len(paths) < 1 {
		return nil, errors.New("Bad path supplied: " + path)
	}
	P, rest := paths[0], paths[1:]

	switch P {
	case "proj":
		return utils.GetByPathSlice(rest, D.Proj)

	case "type":
		return utils.GetByPathSlice(rest, D.Type)

	case "dsl":
		return utils.GetByPathSlice(rest, D.Dsl)

	case "custom":
		return utils.GetByPathSlice(rest, D.Custom)

	default:
		return nil, errors.New("Unknown path start for design: " + P)

	}

}
