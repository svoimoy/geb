{{#with DslContext as |CTX|}}
package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
	log15 "gopkg.in/inconshreveable/log15.v2"

	{{#if CTX.resources}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/resources"
	{{/if}}
	{{#if CTX.routes}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/routes"
	{{/if}}

	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// Name:     {{name}}
// Version:  {{version}}
// About:    {{about}}

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

// main package logger
var logger = log15.New()

func main() {

	// load the configuration file
	read_config()

	// configure the logger
	config_logger()

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
	G := E.Group("{{config.base-url}}")

	setupRoutes(G)

	// HOFSTADTER_START main-prerun
	// HOFSTADTER_END   main-prerun

	host := viper.GetString("host")
	port := viper.GetString("port")

	E.Logger.SetLevel(log.INFO)
	E.Logger.Fatal(E.Start(host+":"+port))
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

func config_logger() {
	// log-config default global values
	level := log15.LvlWarn
	stack := false

	// look up in config
	lcfg := viper.GetStringMap("log-config.default")

	if lcfg != nil && len(lcfg) > 0 {
		level_str := lcfg["level"].(string)
		stack = lcfg["stack"].(bool)
		level_local, err := log15.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}
		level = level_local
	}

	termlog := log15.LvlFilterHandler(level, log15.StdoutHandler)
	if stack {
		term_stack := log15.CallerStackHandler("%+v", log15.StdoutHandler)
		termlog = log15.LvlFilterHandler(level, term_stack)
	}

	logger.SetHandler(termlog)

	// set package loggers
	xtalk.SetLogger(logger)
	{{#if CTX.resources}}
	resources.SetLogger(logger)
	{{/if}}
	{{#if CTX.routes}}
	routes.SetLogger(logger)
	{{/if}}

	// HOFSTADTER_START config-logger
	// HOFSTADTER_END   config-logger
}

{{/with}}
