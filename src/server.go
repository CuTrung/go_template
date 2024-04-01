package server

import (
	"github.com/CuTrung/go_template/src/confs"
	"github.com/CuTrung/go_template/src/db"
	"github.com/CuTrung/go_template/src/middlewares"
	"github.com/CuTrung/go_template/src/modules/book"
	"github.com/CuTrung/go_template/src/routers"
	"github.com/gin-gonic/gin"
)

func InitApp() *gin.Engine {
	r := gin.New()
	r.Static("/assets", "./assets")
	confs.SetConfigs(r)
	middlewares.ApplyMiddlewares(r)
	routers.InitRouters(r)
	db.Migrate(&book.Book{})
	return r
}
