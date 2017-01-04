package engine

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/spf13/viper"
	"github.ibm.com/hofstadter-io/dotpath"
	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/gen"
	"github.ibm.com/hofstadter-io/geb/engine/project"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
	"gopkg.in/yaml.v2"
)

func ViewGeb(args []string) (string, error) {
	fmt.Println("ViewGeb:", args)

	// file := utils.LookForKnownFiles()

	data := viper.AllSettings()

	if len(args) == 0 {
		bytes, err := yaml.Marshal(data)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		return string(bytes), nil
	}

	logger.Info("DATA", "data", data)
	ret := ""
	for _, path := range args {
		ret += fmt.Sprintln("path:    ", path, "\n--------------------------------")
		val, err := dotpath.Get(path, data, true)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}

		bytes, err := yaml.Marshal(val)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}
		ret += string(bytes)
	}

	return ret, nil
}

func ViewDsl(folder string, args []string) (string, error) {
	fmt.Println("ViewDsl:", args)

	file := utils.LookForKnownFiles()

	var data interface{}
	switch file {
	case "geb.yml", "geb.yaml":
		P := project.NewProject()
		err := P.Load(file, nil)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		dsl_map := map[string]interface{}{}
		for key, D := range P.DslMap {
			dsl_map[key] = D.Generators
		}

		data = dsl_map

	case "geb-dsl.yml", "geb-dsl.yaml":
		D, err := dsl.CreateFromFolder(folder)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		data = D

	default:
		return "", errors.New("Default TBD, should load from system library...")
	}

	if len(args) == 0 {
		bytes, err := yaml.Marshal(data)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		return string(bytes), nil
	}

	logger.Info("DATA", "data", data)
	ret := ""
	for _, path := range args {
		ret += fmt.Sprintln("path:    ", path, "\n--------------------------------")
		val, err := dotpath.Get(path, data, true)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}

		bytes, err := yaml.Marshal(val)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}
		ret += string(bytes)
	}

	return ret, nil
}

func ViewGen(folder string, args []string) (string, error) {
	fmt.Println("ViewGen:", args)
	file := utils.LookForKnownFiles()

	var data interface{}
	switch file {
	case "geb.yml", "geb.yaml":
		P := project.NewProject()
		err := P.Load(file, nil)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		gen_map := map[string]interface{}{}
		for key, D := range P.DslMap {
			gen_map[key] = D.Generators
		}

		data = gen_map

	case "geb-dsl.yml", "geb-dsl.yaml":
		D, err := dsl.CreateFromFolder(folder)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		data = D.Generators

	case "geb-gen.yml", "geb-gen.yaml":
		G, err := gen.CreateFromFolder(folder)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		data = G

	default:
		return "", errors.New("Default TBD, should load from system library...")
	}

	if len(args) == 0 {
		bytes, err := yaml.Marshal(data)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		return string(bytes), nil
	}

	ret := ""
	for _, path := range args {
		ret += fmt.Sprintln("path:    ", path, "\n--------------------------------")
		val, err := dotpath.Get(path, data, true)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}

		bytes, err := yaml.Marshal(val)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}
		ret += string(bytes)
	}

	return ret, nil
}

func ViewDesign(folder string, args []string) (string, error) {
	fmt.Println("ViewDesign:", folder, args)

	D, err := design.CreateFromFolder(folder)
	if err != nil {
		return "", errors.Wrap(err, "in engine.ViewDesign")
	}

	if len(args) == 0 {
		bytes, err := yaml.Marshal(D)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewDesign")
		}
		return string(bytes), nil
	}

	ret := ""
	for i, path := range args {
		if path == "" {
			logger.Error("Empty path component in ViewDesign", "i", i, "args", args)
			continue
		}
		ret += fmt.Sprintln("path:    ", path, "\n--------------------------------")
		val, err := D.GetByPath(path)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewDesign")
		}

		bytes, err := yaml.Marshal(val)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewDesign")
		}
		ret += string(bytes)
	}

	return ret, nil
}

func ViewProject(folder string, args []string) (string, error) {
	fmt.Println("ViewProject:", args)

	file := utils.LookForKnownFiles()

	var data interface{}
	switch file {
	case "geb.yml", "geb.yaml":
		P := project.NewProject()
		err := P.Load(file, nil)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		data = P

	default:
		return "", errors.New("Default TBD, should load from system library...")
	}

	if len(args) == 0 {
		bytes, err := yaml.Marshal(data)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		return string(bytes), nil
	}

	ret := ""
	for _, path := range args {
		ret += fmt.Sprintln("path:    ", path, "\n--------------------------------")
		val, err := dotpath.Get(path, data, true)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}

		bytes, err := yaml.Marshal(val)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}
		ret += string(bytes)
	}

	return ret, nil
}

func ViewPlans(folder string, args []string) (string, error) {
	fmt.Println("ViewProject:", args)

	file := utils.LookForKnownFiles()

	var data interface{}
	switch file {
	case "geb.yml", "geb.yaml":
		P := project.NewProject()
		err := P.Load(file, nil)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		err = P.Plan()
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		data = P.Plans

	default:
		return "", errors.New("Default TBD, should load from system library...")
	}

	if len(args) == 0 {
		bytes, err := yaml.Marshal(data)
		if err != nil {
			return "", errors.Wrap(err, "in engine.ViewGen")
		}
		return string(bytes), nil
	}

	ret := ""
	for _, path := range args {
		ret += fmt.Sprintln("path:    ", path, "\n--------------------------------")
		val, err := dotpath.Get(path, data, true)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}

		bytes, err := yaml.Marshal(val)
		if err != nil {
			return ret, errors.Wrap(err, "in engine.ViewGen")
		}
		ret += string(bytes)
	}

	return ret, nil
}
