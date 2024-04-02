package routers

import (
	"os"

	"github.com/CuTrung/go_template/src/modules/book"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	rg := r.Group(os.Getenv("API_PREFIX"))

	book_route := rg.Group("/books")
	{
		book_route.GET("", book.GetBook)
		book_route.POST("", book.CreateBook)
	}
}
