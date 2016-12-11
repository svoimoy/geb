package system

import (
	log "gopkg.in/inconshreveable/log15.v2"
)

var logger log.Logger

func SetLogger(l log.Logger) {
	logger = l
}


/*
Repeated Context
----------------
long: View the geb system configuration
name: View
parent: System
path: commands.subcommands
short: View the geb system configuration
usage: view

*/
