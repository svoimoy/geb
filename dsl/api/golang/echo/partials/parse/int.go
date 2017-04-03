{{name}}_int, err := strconv.ParseInt({{name}}, 10, 64)
if err != nil {
	res := gin.H{"error": "{{name}} must be an unsigned integer"}
	c.JSON(http.StatusBadRequest, res)
	return
}

