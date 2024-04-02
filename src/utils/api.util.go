package utils

import (
	"time"

	"github.com/CuTrung/go_template/src/common/types"
	"github.com/gin-gonic/gin"
)

func Transform[T any](obj *T) func(c *gin.Context) error {
	return func(c *gin.Context) error {
		if err := c.ShouldBindJSON(&obj); err != nil {
			return err
		}
		return nil
	}
}

func JSON[T any](status int, res types.ApiResponse[T]) func(c *gin.Context) {
	return func(c *gin.Context) {
		if IsDevelopment() {
			c.IndentedJSON(status, res)
			return
		}
		c.JSON(status, res)
	}
}

func FormatJSON[T any](status int, data T) func(c *gin.Context) {
	return func(c *gin.Context) {
		res := types.ApiResponse[T]{
			Data: data,
			Date: time.Now(),
			Path: c.Request.URL.Path,
		}

		JSON(status, res)(c)
	}
}

func FormatErrorJSON(status int, errs []string) func(c *gin.Context) {
	return func(c *gin.Context) {
		res := types.ApiResponse[any]{
			Date:   time.Now(),
			Path:   c.Request.URL.Path,
			Errors: errs,
		}

		JSON(status, res)(c)
	}
}
