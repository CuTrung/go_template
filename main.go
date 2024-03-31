package main

import (
	"fmt"
	"os"

	server "github.com/CuTrung/go_template/src"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	server.InitApp(r)
	url := fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
	r.Run(url)
}
