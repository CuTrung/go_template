package confs

import (
	"github.com/gin-gonic/gin"
)

func SetConfigs(r *gin.Engine) {
	SetTrustedProxies(r)
	SetMode()
}
