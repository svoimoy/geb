package main

import (
	"github.com/gin-gonic/gin"
)

// Name:     count-service
// Version:  0.0.1
// About:    AlchemyAPI count-service for daily transaction counts.

func main() {
	router := gin.Default()

	group := router.Group("api/v1")
	group.GET("/version", version_GET_Handler)
	group.GET("/all_counts", all_counts_GET_Handler)
	group.GET("/trans_count", trans_count_GET_Handler)
	group.PUT("/trans_count", trans_count_PUT_Handler)
	group.DELETE("/trans_count", trans_count_DELETE_Handler)

	r.Run() // listen and server on 0.0.0.0:8080
}
