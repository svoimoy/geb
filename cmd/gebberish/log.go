package gebberish

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
args:
- help: The rule to apply [r# or rule-#]
  name: rule
  required: true
  type: string
- help: optional args to rules 3 and 4
  name: extra
  rest: true
  type: array:string
long: |
  Welcome to the MI game

  start with mi-string = 'MI'

  mi-rule-1:    if mi-string ends in 'I',        you may add a 'U'
  mi-rule-2:    suppose mi-string = 'Mx',          then you may make it 'Mxx'
  mi-rule-3:    if mi-string contains an 'III',  you may replace it with 'U'
  mi-rule-4:    if mi-string contains a 'UU',    you may drop it (remove it)

  Goal: Try to get 'MU'

  Input: r1,r2,r3,r4 or reset

  Notes:
   - rules 3 and 4 require a zero-indexed postional argument to determine the occurance to remove
   - use 'reset' to restore mi-string to 'MI'
name: Mi
parent: Gebberish
path: commands.subcommands
short: View information about a Project's Plans
usage: mi

*/
