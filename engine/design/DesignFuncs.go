package design

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Design
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewDesign() *Design {
	return &Design{
		Files:  map[string]interface{}{},
		Proj:   map[string]interface{}{},
		Data:   map[string]interface{}{},
		Type:   map[string]interface{}{},
		Pkg:    map[string]interface{}{},
		Dsl:    map[string]interface{}{},
		Custom: map[string]interface{}{},
		Extra:  map[string]interface{}{},
	}
}

/*
Where's your docs doc?!
*/
func (D *Design) ImportDesignFile(filename string) (err error) {
	// HOFSTADTER_START ImportDesignFile
	logger.Info("Importing Design filename:", "filename", filename)
	return D.importDesign(filepath.Dir(filename), filename)
	// HOFSTADTER_END   ImportDesignFile
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) ImportDesignFolder(folder string) (err error) {
	// HOFSTADTER_START ImportDesignFolder
	logger.Info("Importing Design folder: " + folder)

	// Make sure the folder exists
	_, err = os.Lstat(folder)
	if err != nil {
		return errors.Wrap(err, "in design.ImportDesignFolder: "+folder+"\n")
	}

	// local walk function closure
	import_design_walk_func := func(path string, info os.FileInfo, err error) error {
		local_d := D
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)[1:]
		switch ext {

		case "json", "toml", "xml", "yaml", "yml", "hof":
			lerr := local_d.importDesign(folder, path)
			if lerr != nil {
				logger.Debug("importing error", "path", path, "error", lerr)
				return errors.Wrap(err, "in design.ImportDesignFolder: "+folder+"  "+path+"\n")
			}
			return nil
		default:
			return nil
		}
	}

	// Walk the directory
	err = filepath.Walk(folder, import_design_walk_func)
	if err != nil {
		return errors.Wrap(err, "in design.ImportDesignFolder: "+folder+"\n")
	}
	return nil
	// HOFSTADTER_END   ImportDesignFolder
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) ImportExtraFile(filename string) (err error) {
	// HOFSTADTER_START ImportExtraFile
	logger.Info("Importing Design filename:", "filename", filename)
	return D.importExtra(filepath.Dir(filename), filename)
	// HOFSTADTER_END   ImportExtraFile
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) ImportExtraFolder(folder string) (err error) {
	// HOFSTADTER_START ImportExtraFolder
	logger.Info("Importing Design folder: " + folder)

	// Make sure the folder exists
	_, err = os.Lstat(folder)
	if err != nil {
		return errors.Wrap(err, "in design.ImportDesignFolder: "+folder+"\n")
	}

	// local walk function closure
	import_design_walk_func := func(path string, info os.FileInfo, err error) error {
		local_d := D
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)[1:]
		switch ext {

		case "json", "toml", "xml", "yaml", "yml", "hof":
			lerr := local_d.importExtra(folder, path)
			if lerr != nil {
				logger.Debug("importing error", "path", path, "error", lerr)
				return errors.Wrap(err, "in design.ImportDesignFolder: "+folder+"  "+path+"\n")
			}
			return nil
		default:
			return nil
		}
	}

	// Walk the directory
	err = filepath.Walk(folder, import_design_walk_func)
	if err != nil {
		return errors.Wrap(err, "in design.ImportDesignFolder: "+folder+"\n")
	}
	return nil
	// HOFSTADTER_END   ImportExtraFolder
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) Get(path string) (object interface{}, err error) {
	// HOFSTADTER_START Get
	return D.GetByPath(path)
	// HOFSTADTER_END   Get
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) GetByPath(path string) (object interface{}, err error) {
	// HOFSTADTER_START GetByPath
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

		case "extra":
			return D.Extra, nil

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

	case "extra":
		return dotpath.GetByPathSlice(rest, D.Extra, true)

	default:
		return nil, errors.New("Unknown path start for design: " + P)

	}
	// HOFSTADTER_END   GetByPath
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) Validate() (errorReport map[string]error) {
	// HOFSTADTER_START Validate

	// HOFSTADTER_END   Validate
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) importDesign(basePath string, designPath string) (err error) {
	// HOFSTADTER_START ImportDesign
	// HOFSTADTER_END   ImportDesign
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) storeDesign(relativePath string, dsl string, design interface{}) (err error) {
	// HOFSTADTER_START StoreDesign
	// HOFSTADTER_END   StoreDesign
	return
}

/*
Where's your docs doc?!
*/
func (D *Design) importExtra(basePath string, designPath string) (err error) {
	// HOFSTADTER_START ImportExtra
	// HOFSTADTER_END   ImportExtra
	return
}

// HOFSTADTER_BELOW
