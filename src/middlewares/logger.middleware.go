package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/CuTrung/go_template/src/common/consts"
	"github.com/CuTrung/go_template/src/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func createLogger(file_name string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   file_name,
		MaxSize:    4, // Dung lượng tối đa là 4MB
		MaxBackups: 100,
		MaxAge:     28,   // Số ngày giữ lại các file log cũ
		Compress:   true, // Nén các file log cũ
	}
}
func LoggerMiddleware(c *gin.Context) {
	start := time.Now()
	prefix_log := fmt.Sprintf("logs/%v_%v", utils.GetCurrentDate(), os.Getenv(consts.NODE_ENV))
	info_path, error_path := fmt.Sprintf("%v.info.log", prefix_log), fmt.Sprintf("%v.error.log", prefix_log)
	infoLogger, errorLogger := createLogger(info_path), createLogger(error_path)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, infoLogger, errorLogger)
	c.Next()
	status := c.Writer.Status()
	duration := time.Since(start)
	level_logger := infoLogger
	request_id := c.Request.Header[consts.HEADER_REQUEST_ID][0]
	if status >= 400 {
		level_logger = errorLogger
	}
	format_logger := fmt.Sprintf("[%s] [%s] %d %s %s `%v` (%s)\n", os.Getenv(consts.APP_NAME), utils.GetCurrentDateHasHour(), status, c.Request.Method, c.Request.URL.Path, request_id, duration)
	fmt.Fprint(level_logger, format_logger)
	fmt.Println(format_logger)
}
