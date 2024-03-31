package confs

import "github.com/gin-gonic/gin"

func SetTrustedProxies(r *gin.Engine) {
	white_list := []string{
		"localhost:4000",
	}
	r.ForwardedByClientIP = true
	r.SetTrustedProxies(white_list)
}
