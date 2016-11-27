package gen

import (
	log "gopkg.in/inconshreveable/log15.v2" // logging framework
)

var logger log.Logger

func SetLogger(l log.Logger) {
	logger = l
}
