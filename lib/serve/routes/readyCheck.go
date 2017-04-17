package routes

import (
	"github.com/labstack/echo"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
API:     serve
Name:    ready-check
Route:   readyz
Method:  GET
Path:    routes
Parent:  serve
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

// Should find a way to build up errors and return all
// GET    ->
func Handle_GET_ReadyCheck(ctx echo.Context) (err error) {
	// Check params

	// HOFSTADTER_START handler
	// HOFSTADTER_END   handler

	return nil
}

// HOFSTADTER_BELOW
