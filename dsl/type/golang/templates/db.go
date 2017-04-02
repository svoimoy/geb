{{#with DslContext as |DSL| }}
package {{#each (split DSL.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	// HOFSTADTER_START import
	"github.com/jinzhu/gorm"
	// HOFSTADTER_END   import
)

/*
Name:      {{DSL.name}}
About:     {{DSL.about}}
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

func init() {
	// ???

}

func AutoMigrate{{camelT DSL.name}}(db *gorm.DB) {
	db.Debug().AutoMigrate(&{{camelT DSL.name}}{})
}

{{/with}}

// HOFSTADTER_BELOW
