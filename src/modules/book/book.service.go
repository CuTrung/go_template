package book

import (
	"fmt"
	"net/http"

	"github.com/CuTrung/go_template/src/db"
	"github.com/CuTrung/go_template/src/utils"
	"github.com/gin-gonic/gin"
)

func GetBookService() (int, any) {
	return http.StatusOK, gin.H{
		"message": "book",
	}
}

func CreateBookService(create_book_dto CreateBookDTO) (int, any) {
	var db = db.GetInstance()
	result := db.Create(&create_book_dto) // pass pointer of data to Create

	fmt.Print(">>> error", result.Error, create_book_dto)

	return http.StatusCreated, utils.ToGinObj(create_book_dto)
}
