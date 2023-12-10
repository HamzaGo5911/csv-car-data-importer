package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const defaultLimit = 10

func Paginate(c *gin.Context, totalCount int, data interface{}) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset := (page - 1) * limit

	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"page":   page,
		"limit":  limit,
		"total":  totalCount,
		"offset": offset,
	})
}
