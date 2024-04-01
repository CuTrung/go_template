package middlewares

import (
	"fmt"
	"net/http"

	"github.com/CuTrung/go_template/src/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()
	status := http.StatusBadRequest
	errors := []string{}
	for _, err := range c.Errors {
		errors = append(errors, err.Error())
	}

	status_res := c.Writer.Status()
	if len(errors) <= 0 && status_res > 400 {
		status = http.StatusNotFound
		errors = append(errors, fmt.Sprintf("Path %v not found", c.Request.URL.Path))
	}

	if len(errors) > 0 {
		utils.FormatJSON(status, gin.H{
			"errors": errors,
		})(c)
	}
}
