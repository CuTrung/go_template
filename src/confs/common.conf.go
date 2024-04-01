package confs

import (
	"github.com/CuTrung/go_template/src/common/consts"
	"github.com/CuTrung/go_template/src/utils"
	"github.com/gin-gonic/gin"
)

func SetTrustedProxies(r *gin.Engine) {
	white_list := []string{
		"localhost:4000",
	}
	r.ForwardedByClientIP = true
	r.SetTrustedProxies(white_list)
}

func SetMode() {
	isProduction := utils.IsProduction()
	mode := consts.DEBUG_MODE
	if isProduction {
		mode = consts.RELEASE_MODE
	}
	gin.SetMode(mode)
}
