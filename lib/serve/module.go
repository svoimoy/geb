package serve

import (
	"os"
	"sort"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"

	"github.ibm.com/hofstadter-io/geb/lib/serve/resources"

	"github.ibm.com/hofstadter-io/geb/lib/serve/routes"

	// HOFSTADTER_START import
	"fmt"
	// HOFSTADTER_END   import
)

// Name:     serve
// Version:
// About:    serve templates and ETL pipelines with geb

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
type Routes []echo.Route

func (slice Routes) Len() int {
    return len(slice)
}

func (slice Routes) Less(i, j int) bool {
    if slice[i].Path == slice[j].Path {
    	return slice[i].Method < slice[j].Method
	}
    return slice[i].Path < slice[j].Path
}

func (slice Routes) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}
// HOFSTADTER_END   init

func Run() (err error) {

	// load the configuration file
	read_config()

	// create the echo server object
	E := echo.New()

	// Pre-Middleware
	// HOFSTADTER_START main-pre-middleware
	// HOFSTADTER_END   main-pre-middleware

	// Use-Middleware
	E.Use(middleware.Recover())

	// HOFSTADTER_START main-pre-routes
	// HOFSTADTER_END   main-pre-routes

	// Base API Group
	G := E.Group("/api/v1")

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
	sort.Sort(Routes(routes))
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
