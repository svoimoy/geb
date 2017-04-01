{{#with DslContext as |API| }}
package databases

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"

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

func ConnectToPostgres() {

	host := viper.GetString("pg-host")
	user := viper.GetString("pg-user")
	pass := viper.GetString("pg-pass")
	db := viper.GetString("pg-db")
	sslmode := viper.GetString("pg-sslmode")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, pass, db, sslmode)

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
