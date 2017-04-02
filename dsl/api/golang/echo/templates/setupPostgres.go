{{#with DslContext as |API| }}
package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var POSTGRES *gorm.DB

func ConnectToPostgres(connStr string) {

	var err error
	POSTGRES, err = gorm.Open("postgres", connStr)
	if err != nil {
		logger.Error("Unable to connect to Postgres", "error", err, "connStr", connStr)
	}
}

func DisconnectFromPostgres() {
	POSTGRES.Close()
}

{{/with}}
// HOFSTADTER_BELOW
