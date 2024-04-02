package confs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetConfigs(r *gin.Engine) {
	SetTrustedProxies(r)
	SetMode()
	db = ConnectDB()
}

func GetDBInstance() *gorm.DB {
	return db
}
