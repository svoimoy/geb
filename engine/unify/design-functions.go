package unify

import (
	"io/ioutil"
)

type designFunction func(string) (interface{}, error)

var design_functions = map[string]designFunction{
	"loadfile": loadFile,
}

func loadFile(arg string) (interface{}, error) {

	content, err := ioutil.ReadFile(arg)
	if err != nil {
		return err.Error(), err
	}

	return string(content), nil
}
