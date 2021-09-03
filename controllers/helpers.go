package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination helps in paginating the response
func Pagination(c *gin.Context) (pageNumberValue int, pageSizeValue int, err error) {
	pageNumber := c.Request.URL.Query().Get("pageNumber")
	if pageNumber == "" {
		pageNumberValue = DefaultPageNumber
	} else {
		number, er := strconv.Atoi(pageNumber)
		if er == nil {
			pageNumberValue = number
		} else {
			return pageNumberValue, pageSizeValue, errors.New("invalid pageNumber: " + pageNumber)
		}
	}

	pageSize := c.Request.URL.Query().Get("pageSize")
	if pageSize == "" {
		pageSizeValue = DefaultPageSize
	} else {
		number, er := strconv.Atoi(pageSize)
		if er == nil {
			pageSizeValue = number
		} else {
			return pageNumberValue, pageSizeValue, errors.New("invalid pageSize: " + pageSize)
		}
	}

	return

}
