package serve

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"

	"github.com/hofstadter-io/geb/lib/serve/resources"

	"github.com/hofstadter-io/geb/lib/serve/routes"

	// HOFSTADTER_START import
	"fmt"
	"sort"
	// HOFSTADTER_END   import
)

// Name:     serve
// Version:
// About:    serve templates and ETL pipelines with geb

// HOFSTADTER_START start
// HOFSTADTER_END   start

func Run() (err error) {

	// load the configuration file
	read_config()

	// create the echo server object
	E := echo.New()

	// Use-Middleware
	// E.Use(middleware.Recover())

	// HOFSTADTER_START main-pre-group
	// HOFSTADTER_END   main-pre-group

	// Base API Group
	G := E.Group("/api/v1")

	// HOFSTADTER_START main-pre-middleware
	// HOFSTADTER_END   main-pre-middleware

	AddMiddleware(G, []string{"default"})

	// HOFSTADTER_START main-pre-routes
	// HOFSTADTER_END   main-pre-routes

	err = resources.InitRouter(G)
	if err != nil {
		return err
	}

	err = routes.InitRouter(G)
	if err != nil {
		return err
	}

	// HOFSTADTER_START main-prerun
	fmt.Println("Routes:")

	routes := E.Routes()
	sort.Slice(routes, func(i, j int) bool {
		if routes[i].Path == routes[j].Path {
			return routes[i].Method < routes[j].Method
		}
		return routes[i].Path < routes[j].Path
	})

	for _, r := range routes {
		fmt.Printf("  %8s:  %s\n", r.Method, r.Path)
	}
	// HOFSTADTER_END   main-prerun

	appHost := viper.GetString("host")
	appPort := viper.GetString("port")

	E.Logger.SetLevel(log.INFO)
	err = E.Start(appHost + ":" + appPort)

	if err != nil {
		return err
	}

	return nil
}

func read_config() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.MergeInConfig()

	// Hackery because viper only takes the first config file found... not merging, wtf does merge config mean then anyway
	f, err := os.Open("config.yml")
	if err != nil {
		f = nil
		f2, err2 := os.Open("config.yaml")
		if err2 != nil {
			f = nil
		} else {
			f = f2
		}
	}
	if f != nil {
		verr := viper.MergeConfig(f)
		if verr != nil {
			panic(verr)
		}
	} else {
		panic("missing config.yaml during start up")
	}
}

// HOFSTADTER_BELOW
