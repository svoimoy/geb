package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func version_GET_Handler(ctx *gin.Context) {
}

func all_counts_GET_Handler(ctx *gin.Context) {
}

func trans_count_GET_Handler(ctx *gin.Context) {
	owner_id := c.Query("owner_id")
	if owner_id == "" {
		res := gin.H{"error": "missing owner_id in request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// not a uint: 'string'
}

func trans_count_PUT_Handler(ctx *gin.Context) {
	owner_id := c.Query("owner_id")
	if owner_id == "" {
		res := gin.H{"error": "missing owner_id in request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// not a uint: 'string'
	count := c.Query("count")
	if count == "" {
		res := gin.H{"error": "missing count in request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	count_int, err := strconv.ParseUint(count, 10, 64)
	if err != nil {
		res := gin.H{"error": "count must be an unsigned integer"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
}

func trans_count_DELETE_Handler(ctx *gin.Context) {
	date := c.Query("date")
	if date == "" {
		res := gin.H{"error": "missing date in request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// not a uint: 'string'
}
