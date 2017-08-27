package routes

import (
	"github.com/labstack/echo"
	"net/http"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START start
// HOFSTADTER_END   start

var (
	ReadyStatus = false
	ReadyState  = "NotReady"
	ReadyInfo   = map[string]string{
		"reason":  "unknown",
		"message": "unchanged since startup",
	}
)

func addKubernetesHandlers(G *echo.Group) (err error) {

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	group := G.Group("")

	// HOFSTADTER_START router-mid
	// HOFSTADTER_END   router-mid

	group.GET("/lively", Handle_GET_LivelyCheck)
	group.GET("/ready", Handle_GET_ReadyCheck)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

func Handle_GET_LivelyCheck(ctx echo.Context) (err error) {

	// return object
	ret := map[string]interface{}{
		"status": "alive",
	}

	// HOFSTADTER_START handler
	// HOFSTADTER_END   handler

	return ctx.JSON(http.StatusOK, ret)
}

func Handle_GET_ReadyCheck(ctx echo.Context) (err error) {

	// return object
	ret := map[string]interface{}{
		"status": ReadyStatus,
		"state":  ReadyState,
		"info":   ReadyInfo,
	}

	// return code
	code := http.StatusOK
	if !ReadyStatus {
		code = http.StatusServiceUnavailable
	}

	// HOFSTADTER_START ready-handler
	// HOFSTADTER_END   ready-handler

	return ctx.JSON(code, ret)
}

// HOFSTADTER_BELOW
