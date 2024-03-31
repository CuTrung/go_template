package book

import (
	"net/http"

	utils "github.com/CuTrung/go_template/src/utils/api"
	"github.com/gin-gonic/gin"
)

func GetBookService() (int, any) {
	return http.StatusOK, gin.H{
		"message": "book",
	}
}

func CreateBookService(book CreateBookDTO) (int, any) {
	return http.StatusCreated, utils.ToGinObj(book)
}
