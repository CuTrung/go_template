package server

import (
	"log"

	"github.com/CuTrung/go_template/src/confs"
	"github.com/CuTrung/go_template/src/middlewares"
	"github.com/CuTrung/go_template/src/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp(r *gin.Engine) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r.Static("/assets", "./assets")
	confs.SetTrustedProxies(r)
	confs.SetMode()
	middlewares.UseMiddlewares(r)
	routers.InitRouters(r)
}
