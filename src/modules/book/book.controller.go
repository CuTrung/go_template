package book

import (
	"github.com/CuTrung/go_template/src/utils"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	utils.FormatJSON(GetBookService())(c)
}

func CreateBook(c *gin.Context) {
	var newBook CreateBookDTO
	err := utils.Transform(&newBook)(c)
	if err != nil {
		c.Error(err)
		return
	}
	utils.FormatJSON(CreateBookService(newBook))(c)
}
