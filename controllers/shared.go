package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getPagingParams(c *gin.Context) (page, pageSize int) {
	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 5
	}

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	return
}
