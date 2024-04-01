package server

import (
	"github.com/CuTrung/go_template/src/confs"
	"github.com/CuTrung/go_template/src/middlewares"
	"github.com/CuTrung/go_template/src/routers"
	"github.com/gin-gonic/gin"
)

func InitApp() *gin.Engine {
	r := gin.New()
	r.Static("/assets", "./assets")
	confs.SetTrustedProxies(r)
	confs.SetMode()
	middlewares.UseMiddlewares(r)
	routers.InitRouters(r)
	return r
}
