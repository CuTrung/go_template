package types

import (
	"time"
)

type ApiResponse[T any] struct {
	Data   T         `json:"data"`
	Errors []string  `json:"errors"`
	Date   time.Time `json:"date"`
	Path   string    `json:"path"`
}

type BaseModel struct {
	// ref to gorm.Model
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
