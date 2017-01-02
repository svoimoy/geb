{{#with dsl.cli as |CLI| }}
package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/' )}}}/cmd"

	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

var logger = log.New()

func main() {
	read_config()
	config_logger()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func read_config() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("{{CLI.name}}")
	viper.AddConfigPath("$HOME/.{{CLI.name}}")
	viper.MergeInConfig()
	viper.AddConfigPath(".")
	viper.MergeInConfig()

	cfg := viper.GetString("add-config")
	if cfg != "" {
		viper.SetConfigFile(cfg)
		viper.MergeInConfig()
	}

	cfg = viper.GetString("set-config")
	if cfg != "" {
		viper.SetConfigFile(cfg)
		viper.ReadInConfig()
	}
}

func config_logger() {
	// log-config default global values
	level := log.LvlWarn
	stack := false

	// look up in config
	lcfg := viper.GetStringMap("log-config.default")

	if lcfg != nil && len(lcfg) > 0 {
		level_str := lcfg["level"].(string)
		stack = lcfg["stack"].(bool)
		level_local, err := log.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}
		level = level_local
	}

	termlog := log.LvlFilterHandler(level, log.StdoutHandler)
	if stack {
		term_stack := log.CallerStackHandler("%+v", log.StdoutHandler)
		termlog = log.LvlFilterHandler(level, term_stack)
	}

	logger.SetHandler(termlog)

	// set package loggers
	cmd.SetLogger(logger)

	// IF using geb-engine dsl
	// dotpath.SetLogger(logger)
	// engine.SetLogger(logger)

	// HOFSTADTER_START config-logger
	// HOFSTADTER_END   config-logger
}

{{/with}}

// HOFSTADTER_BELOW
