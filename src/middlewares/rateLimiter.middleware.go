package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CuTrung/go_template/src/utils"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	utils.FormatErrorJSON(http.StatusTooManyRequests, []string{"Too many requests. Try again in " + time.Until(info.ResetTime).String()})(c)
}

func RateLimiter() gin.HandlerFunc {
	limit, err := strconv.Atoi(os.Getenv("RATE_LIMIT_OF_MINUTE"))
	if err != nil {
		fmt.Println("Error convert rate limit to number", err)
		return func(ctx *gin.Context) {}
	}
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: uint(limit),
	})
	return ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
}
