package book

import "github.com/CuTrung/go_template/src/common/types"

type Book struct {
	types.BaseModel
	Title string  `json:"title" binding:"required"`
	Label *string `json:"label"`
}
