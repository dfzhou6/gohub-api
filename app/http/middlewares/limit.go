package middlewares

import (
	"gohub/pkg/app"
	"gohub/pkg/limiter"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func limitHandler(c *gin.Context, key string, limit string) bool {
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	if rate.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁",
		})
		return false
	}

	return true
}

func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		c.Set("limiter-once", false)
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}
