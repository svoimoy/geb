{{#with RepeatedContext as |RC| }}
{{#with dsl.cli as |CLI| }}
package {{lower RC.parent}}

import (
	log "gopkg.in/inconshreveable/log15.v2"
)

var logger log.Logger

func SetLogger(l log.Logger) {
	logger = l
}

{{/with}}
{{/with}}

/*
Repeated Context
----------------
{{{yaml RepeatedContext}}}
*/
