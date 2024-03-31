package book

import (
	"net/http"

	utils "github.com/CuTrung/go_template/src/utils/api"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	c.IndentedJSON(GetBookService())
}

func CreateBook(c *gin.Context) {
	var newBook CreateBookDTO
	err := utils.Transform(&newBook)(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(CreateBookService(newBook))
}
