package book

import "github.com/gin-gonic/gin"

func GetBook(c *gin.Context) {
	GetBookController(c)
}
