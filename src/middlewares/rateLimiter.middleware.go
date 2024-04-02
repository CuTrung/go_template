package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CuTrung/go_template/src/common/consts"
	"github.com/CuTrung/go_template/src/utils"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func errorHandler(c *gin.Context, info ratelimit.Info) {
	utils.FormatErrorJSON(http.StatusTooManyRequests, []string{"Too many requests. Try again in " + time.Until(info.ResetTime).String()})(c)
}

func initStore(store_type string) (ratelimit.Store, error) {
	limit, err := utils.ToInt(os.Getenv(consts.RATE_LIMIT_OF_MINUTE))
	if err != nil {
		fmt.Println("Error convert rate limit to number", err)
		return nil, err
	}
	switch store_type {
	case consts.REDIS_STORE:
		return ratelimit.RedisStore(&ratelimit.RedisOptions{
			RedisClient: redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%v:%v", os.Getenv(consts.REDIS_STORE_HOST), os.Getenv(consts.REDIS_STORE_PORT)),
			}),
			Rate:  time.Minute,
			Limit: uint(limit),
		}), nil
	default:
		return ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
			Rate:  time.Minute,
			Limit: uint(limit),
		}), nil
	}
}

func RateLimiter() gin.HandlerFunc {
	store, err := initStore(consts.REDIS_STORE)
	if err != nil {
		return func(ctx *gin.Context) {}
	}
	return ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc: func(c *gin.Context) string {
			return c.ClientIP()
		},
	})
}
