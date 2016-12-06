package dsl

import (
	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
)

func (D *Dsl) Validate() []error {
	ret := []error{}

	errs := D.validate_config()
	ret = utils.AccumErrs(ret, errs)

	return ret
}

func (D *Dsl) validate_config() []error {
	ret := []error{}

	if D.Config == nil {
		return []error{errors.New("nil config")}
	}

	errs := D.validate_name()
	ret = utils.AccumErrs(ret, errs)

	errs = D.validate_version()
	ret = utils.AccumErrs(ret, errs)

	return ret
}

func (D *Dsl) validate_name() []error {
	ret := []error{}
	if D.Config.Name == "" {
		err := errors.New("missing name")
		ret = append(ret, err)
	}
	return ret
}

func (D *Dsl) validate_version() []error {
	ret := []error{}
	if D.Config.Version == "" {
		err := errors.New("missing version")
		ret = append(ret, err)
	}
	return ret
}
