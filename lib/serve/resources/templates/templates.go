package templates

import (
	"github.com/pkg/errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"

	"github.com/hofstadter-io/geb/lib/templates"

	// HOFSTADTER_START import
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/spf13/viper"

	T "github.com/hofstadter-io/geb/engine/templates"
	"github.com/hofstadter-io/geb/lib/templates"
	// HOFSTADTER_END   import
)

/*
API:       serve
Name:      templates
Route:     templates
Resource:  type.lib.templates.Template
Path:      resources
Parent:    serve

*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var templatesValidate = validator.New()

/*
Where's your docs doc?!
*/
func Handle_LIST_Templates(ctx echo.Context) (err error) {

	// HOFSTADTER_START list-body-start
	// HOFSTADTER_END   list-body-start

	// OUTPUT
	// user-defined
	// it's a view!
	var ts []templates.Short
	// fmt.Println("list templates")

	// HOFSTADTER_START list-body-mid
	tDir := viper.GetString("template-dir")
	fis, err := ioutil.ReadDir(tDir)
	if err != nil {
		return err
	}

	for _, f := range fis {
		id := f.Name()
		id = strings.TrimSuffix(id, ".json")
		t := templates.TemplateViewShort{
			ID: id,
		}
		ts = append(ts, t)
	}
	// HOFSTADTER_END   list-body-mid

	// HOFSTADTER_START list-body
	// HOFSTADTER_END   list-body

	// HOFSTADTER_START list-body-end
	// HOFSTADTER_END   list-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, ts)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_POST_Templates(ctx echo.Context) (err error) {

	// HOFSTADTER_START post-body-start
	// HOFSTADTER_END   post-body-start

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's a view!
	var inTpl templates.Create
	err = ctx.Bind(&inTpl)
	if err != nil {
		return err
	}
	err = templatesValidate.Struct(inTpl)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors": err,
				"type":   "invalid",
			})
		}
		if _, ok := err.(*validator.ValidationErrors); ok {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors": err,
				"type":   "validation",
			})
		}
		return err
	}
	// END binding input to query/form/body params

	// OUTPUT
	// user-defined
	// it's not a view
	var outTpl templates.Template
	// fmt.Println("post templates")

	// HOFSTADTER_START post-body-mid
	outTpl.ID = uuid.New().String()
	outTpl.Name = inTpl.Name
	outTpl.Data = inTpl.Data

	err = writeTemplate(outTpl)
	if err != nil {
		return err
	}

	// HOFSTADTER_END   post-body-mid

	// HOFSTADTER_START post-body
	// HOFSTADTER_END   post-body

	// HOFSTADTER_START post-body-end
	// HOFSTADTER_END   post-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, outTpl)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_GET_Templates(ctx echo.Context) (err error) {

	// HOFSTADTER_START get-body-start
	// HOFSTADTER_END   get-body-start

	// M path params
	templateId := ctx.Param("templateId")

	// OUTPUT
	// user-defined
	// it's not a view
	var t templates.Template
	// fmt.Println("get templates")

	// HOFSTADTER_START get-body-mid
	T, err := readTemplate(templateID)
	if err != nil {
		return err
	}
	t = T
	// HOFSTADTER_END   get-body-mid

	// HOFSTADTER_START get-body
	// HOFSTADTER_END   get-body

	// HOFSTADTER_START get-body-end
	// HOFSTADTER_END   get-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, t)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_PUT_Templates(ctx echo.Context) (err error) {

	// HOFSTADTER_START put-body-start
	// HOFSTADTER_END   put-body-start

	// M path params
	templateId := ctx.Param("templateId")

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's not a view
	var inTpl templates.Template
	err = ctx.Bind(&inTpl)
	if err != nil {
		return err
	}
	err = templatesValidate.Struct(inTpl)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors": err,
				"type":   "invalid",
			})
		}
		if _, ok := err.(*validator.ValidationErrors); ok {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors": err,
				"type":   "validation",
			})
		}
		return err
	}
	// END binding input to query/form/body params

	// OUTPUT
	// user-defined
	// it's not a view
	var outTpl templates.Template
	// fmt.Println("put templates")

	// HOFSTADTER_START put-body-mid
	// should check for existance

	inTpl.ID = templateID
	outTpl = inTpl

	err = writeTemplate(outTpl)
	if err != nil {
		return err
	}

	// HOFSTADTER_END   put-body-mid

	// HOFSTADTER_START put-body
	// HOFSTADTER_END   put-body

	// HOFSTADTER_START put-body-end
	// HOFSTADTER_END   put-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, outTpl)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_DELETE_Templates(ctx echo.Context) (err error) {

	// HOFSTADTER_START delete-body-start
	// HOFSTADTER_END   delete-body-start

	// M path params
	templateId := ctx.Param("templateId")

	// OUTPUT
	// user-defined
	// it's a view!
	var outTpl templates.Short
	// fmt.Println("delete templates")

	// HOFSTADTER_START delete-body-mid
	tDir := viper.GetString("template-dir")
	filename := filepath.Join(tDir, templateID+".json")

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	outTpl.ID = templateID
	outTpl.Name = "deleted"
	// HOFSTADTER_END   delete-body-mid

	// HOFSTADTER_START delete-body
	// HOFSTADTER_END   delete-body

	// HOFSTADTER_START delete-body-end
	// HOFSTADTER_END   delete-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, outTpl)
	return err // hacky...
}

// End resource.methods

// Should find a way to build up errors and return all
func Handle_POST_Render(ctx echo.Context) (err error) {

	// Check params

	// input
	// path params
	templateId := ctx.Param("templateId")

	// OUTPUT
	// user-defined
	// it's not a view
	var renderReturn map[string]interface{}

	// HOFSTADTER_START templates-render-body
	t, err := readTemplate(templateID)
	if err != nil {
		return err
	}

	body := ctx.Request().Body
	inBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	var inData interface{}
	err = json.Unmarshal(inBytes, &inData)
	if err != nil {
		return err
	}

	tpl, err := raymond.Parse(t.Data)
	if err != nil {
		return err
	}
	Tpl := &T.Template{tpl}
	T.AddHelpersToRaymond(Tpl)

	outputData, err := tpl.Exec(inData)
	if err != nil {
		return err
	}

	r := renderReturn
	r["id"] = t.ID
	r["name"] = t.Name
	r["data"] = t.Data
	r["input"] = inData
	r["output"] = outputData

	return ctx.JSON(http.StatusOK, r)

	// HOFSTADTER_END   templates-render-body

	// should check accept-type here
	return ctx.JSON(http.StatusOK, renderReturn)

	return nil
}

// end resource.routes

// HOFSTADTER_BELOW

func readTemplate(id string) (t templates.Template, err error) {
	tDir := viper.GetString("template-dir")
	filename := filepath.Join(tDir, id+".json")
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return t, errors.Wrap(err, "while reading template json file")
	}
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return t, errors.Wrap(err, "while unmarshalling json")
	}

	return t, nil
}

func writeTemplate(t templates.Template) (err error) {
	tDir := viper.GetString("template-dir")
	filename := filepath.Join(tDir, t.ID+".json")

	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return errors.Wrap(err, "while marshalling json")
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.Wrap(err, "while writing template json file")
	}

	return nil
}
