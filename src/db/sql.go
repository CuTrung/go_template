package db

import (
	"github.com/CuTrung/go_template/src/confs"
	"gorm.io/gorm"
)

func GetInstance() *gorm.DB {
	db := confs.GetDBInstance()
	return db
}

func Migrate[T any](tables ...*T) {
	db := GetInstance()
	for _, table := range tables {
		isExists := db.Migrator().HasTable(&table)
		if !isExists {
			db.Migrator().CreateTable(&table)
		}
	}
}
