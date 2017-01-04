package design

import (
	"github.com/pkg/errors"
	"strings"

	"github.ibm.com/hofstadter-io/dotpath"
)

func (D *Design) Get(path string) (interface{}, error) {
	return D.GetByPath(path)
}

func (D *Design) GetByPath(path string) (interface{}, error) {
	paths := strings.Split(path, ".")
	if len(paths) < 1 {
		return nil, errors.New("Bad path supplied: " + path)
	}

	if len(paths) == 1 {
		switch paths[0] {
		case "custom":
			return D.Custom, nil

		case "dsl":
			return D.Dsl, nil

		case "pkg":
			return D.Pkg, nil

		case "proj":
			return D.Proj, nil

		case "type":
			return D.Type, nil

		default:
			return nil, errors.New("Unknown path start for design: " + paths[0])

		}
	}
	P, rest := paths[0], paths[1:]

	switch P {
	case "custom":
		return dotpath.GetByPathSlice(rest, D.Custom, true)

	case "dsl":
		return dotpath.GetByPathSlice(rest, D.Dsl, true)

	case "pkg":
		return dotpath.GetByPathSlice(rest, D.Pkg, true)

	case "proj":
		return dotpath.GetByPathSlice(rest, D.Proj, true)

	case "type":
		return dotpath.GetByPathSlice(rest, D.Type, true)

	default:
		return nil, errors.New("Unknown path start for design: " + P)

	}

}
