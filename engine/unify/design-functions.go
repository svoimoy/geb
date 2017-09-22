package unify

import (
	"io/ioutil"
	"strings"
)

type designFunction func(string) (interface{}, error)

var design_functions = map[string]designFunction{
	"loadfile": loadFile,
}

func loadFile(arg string) (interface{}, error) {

	data, err := ioutil.ReadFile(arg)
	if err != nil {
		return err.Error(), err
	}
	content := string(data)
	lines := strings.Split(content, "\n")

	cleanContent := ""
	for _, line := range lines {
		if strings.Contains(line, "HOF"+"STADTER"+"_BELOW") {
			break
		}
		cleanContent += line + "\n"
	}
	return string(cleanContent), nil
}
