package book

import "github.com/gin-gonic/gin"

func GetBookController(c *gin.Context) {
	GetBookService(c)
}
