package book

import (
	"github.com/gin-gonic/gin"
)

func GetBookService(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "book",
	})
}
