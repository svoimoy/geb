package serve

import (
	"github.com/pkg/errors"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// Name:     serve
// Version:
// About:    serve templates and ETL pipelines with geb

// HOFSTADTER_START start
// HOFSTADTER_END   start

var (
	DefaultMiddleware = []string{
		"logger",
		"recover",
		"request-id",
		// "remove-trailing-slash",
	}

	RegisteredMiddleware map[string]func(next echo.HandlerFunc) echo.HandlerFunc
)

func init() {

	RegisteredMiddleware = make(map[string]func(echo.HandlerFunc) echo.HandlerFunc)

	RegisteredMiddleware["recover"] = middleware.Recover()
	RegisteredMiddleware["logger"] = middleware.Logger()
	RegisteredMiddleware["request-id"] = middleware.RequestID()
	/*
		RegisteredMiddleware["remove-trailing-slash"] = middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
			  RedirectCode: http.StatusMovedPermanently,
			})
	*/
}

func AddMiddleware(G *echo.Group, middlewares []string) (err error) {
	if len(middlewares) < 1 {
		return errors.New("AddMiddleware: must specify middlewares and/or 'default'\n")
	}

	// look for default and insert
	for i, ware := range middlewares {
		if ware == "default" {
			first := append(middlewares[:i], DefaultMiddleware...)
			after := middlewares[i+1:]
			combined := append(first, after...)
			middlewares = combined
			break
		}
	}

	// add the middlewares
	for _, ware := range middlewares {
		err = addMiddleware(G, ware)
		if err != nil {
			return errors.Wrap(err, "while adding middleware: "+ware+"\n")
		}
	}

	return nil
}

func addMiddleware(G *echo.Group, ware string) (err error) {

	wareFunc, ok := RegisteredMiddleware[ware]
	if !ok {
		return errors.New("Unknown middleware: " + ware)
	}

	G.Use(wareFunc)

	return nil
}

// HOFSTADTER_BELOW
