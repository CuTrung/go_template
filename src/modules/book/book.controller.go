package book

import (
	"net/http"

	"github.com/CuTrung/go_template/src/utils"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	utils.FormatJSON(GetBookService())(c)
}

func CreateBook(c *gin.Context) {
	var newBook Book
	err := utils.Transform(&newBook)(c)
	if err != nil {
		c.Error(err)
		return
	}

	data, err_create := CreateBookService(newBook)
	if err_create != nil {
		c.Error(err_create)
		return
	}

	utils.FormatJSON(http.StatusCreated, data)(c)
}
