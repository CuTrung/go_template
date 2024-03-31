package utils

import (
	"reflect"

	"github.com/CuTrung/go_template/src/utils"
	"github.com/gin-gonic/gin"
)

func Transform[T any](obj *T) func(c *gin.Context) error {
	return func(c *gin.Context) error {
		if err := c.ShouldBindJSON(&obj); err != nil {
			return err
		}
		return nil
	}
}

func FormatResponse(res Response) any {
	if res.date_time == "" {
		res.date_time = utils.GetCurrentDate()
	}
	return res
}

func ToGinObj[T any](data T) gin.H {
	result := make(gin.H)

	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i).Interface()
		fieldName := typ.Field(i).Name
		result[fieldName] = field
	}

	return result
}
