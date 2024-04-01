package book

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Label *string
}

type CreateBookDTO = Book

func GetBookModel() *Book {
	return &Book{}
}
