package middlewares

import (
	"github.com/gin-gonic/gin"
)

func UseMiddlewares(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(LoggerMiddleware)
	r.Use(ErrorHandlerMiddleware)
}
