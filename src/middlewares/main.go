package middlewares

import (
	"github.com/gin-gonic/gin"
)

func UseMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())
	// r.Use(LoggerMiddleware())

	// infoLogger, errorLogger := createLogger("info.log"), createLogger("error.log")
	// gin.DefaultWriter = io.MultiWriter(os.Stdout, errorLogger)
	// r.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: io.MultiWriter(os.Stdout, infoLogger)}))
	// r.Use(LoggerMiddlewareV2(infoLogger, errorLogger))

	r.Use(LoggerMiddleware())
}
