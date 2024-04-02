package book

import (
	"net/http"

	"github.com/CuTrung/go_template/src/db"
	"github.com/gin-gonic/gin"
)

func GetBookService() (int, any) {
	return http.StatusOK, gin.H{
		"message": "book",
	}
}

func CreateBookService(book Book) (Book, error) {
	var db = db.GetInstance()
	result := db.Create(&book)

	if result.Error != nil {
		return Book{}, result.Error
	}

	return book, nil
}
