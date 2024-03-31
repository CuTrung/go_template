package middlewares

import (
	"fmt"
	"io"
	"os"

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

func LoggerMiddleware() gin.HandlerFunc {
	prefix_log := fmt.Sprintf("logs/%v", utils.GetCurrentDate())
	info_path, error_path := fmt.Sprintf("%v.info.log", prefix_log), fmt.Sprintf("%v.error.log", prefix_log)
	infoLogger, errorLogger := createLogger(info_path), createLogger(error_path)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, infoLogger, errorLogger)

	return func(c *gin.Context) {
		c.Next()

		// Lấy status code của yêu cầu
		status := c.Writer.Status()

		// Kiểm tra nếu status code lớn hơn hoặc bằng 500
		if status >= 400 {
			// Ghi log vào error.log
			fmt.Fprintf(errorLogger, "[%d] %s %s\n", status, c.Request.Method, c.Request.URL.Path)
		} else {
			// Ghi log vào info.log
			fmt.Fprintf(infoLogger, "[%d] %s %s\n", status, c.Request.Method, c.Request.URL.Path)
		}
		fmt.Printf("[%d] %s %s\n", status, c.Request.Method, c.Request.URL.Path)
	}
}
