package main

import (
	"fmt"

	"github.com/CuTrung/go_template/confs"
	"github.com/CuTrung/go_template/modules/book"
	"github.com/gin-gonic/gin"
)

func main() {
	envs := confs.LoadEnv()
	r := gin.Default()
	r.GET("/books", book.GetBook)
	url := fmt.Sprintf("%v:%v", envs["HOST"], envs["PORT"])
	r.Run(url) // listen and serve on 0.0.0.0:8080
}
