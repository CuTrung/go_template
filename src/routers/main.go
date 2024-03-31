package routers

import (
	"github.com/CuTrung/go_template/src/modules/book"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	v1 := r.Group("/books")
	{
		v1.GET("", book.GetBook)
		v1.POST("", book.CreateBook)
	}
}
