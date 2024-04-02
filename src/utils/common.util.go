package utils

import (
	"os"

	"github.com/CuTrung/go_template/src/common/consts"
)

func IsDevelopment() bool {
	return os.Getenv(consts.NODE_ENV) == consts.DEVELOPMENT
}

func IsProduction() bool {
	return os.Getenv(consts.NODE_ENV) == consts.PRODUCTION
}
