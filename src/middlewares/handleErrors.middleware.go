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

	if len(errors) <= 0 {
		status = http.StatusNotFound
		errors = append(errors, fmt.Sprintf("Path %v not found", c.Request.URL.Path))
	}

	utils.FormatJSON(status, gin.H{
		"errors": errors,
	})(c)
}
