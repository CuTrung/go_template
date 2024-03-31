package book

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title" binding:"required"`
}

type CreateBookDTO = Book
