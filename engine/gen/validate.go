package gen

import (
	"github.com/pkg/errors"

	"github.com/hofstadter-io/geb/engine/utils"
)

func (G *Generator) Validate() []error {
	ret := []error{}

	errs := G.validate_config()
	ret = utils.AccumErrs(ret, errs)

	return ret
}

func (G *Generator) validate_config() []error {
	ret := []error{}

	errs := G.validate_name()
	ret = utils.AccumErrs(ret, errs)

	errs = G.validate_version()
	ret = utils.AccumErrs(ret, errs)

	return ret
}

func (G *Generator) validate_name() []error {
	errs := []error{}
	if G.Config.Name == "" {
		err := errors.New("missing name")
		errs = append(errs, err)
	}
	return errs
}

func (G *Generator) validate_version() []error {
	errs := []error{}
	if G.Config.Version == "" {
		err := errors.New("missing version")
		errs = append(errs, err)
	}
	return errs
}

func (G *Generator) validate_repeated() []error {
	errs := []error{}
	if len(G.Repeated) == 0 {
		return errs
	}

	return errs
}
