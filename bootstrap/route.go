package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(route *gin.Engine) {
	// 设置路由
	registerGlobalMiddleware(route)

	routes.RegisterAPIRoutes(route)

	setUp404Handler(route)
}

func registerGlobalMiddleware(route *gin.Engine) {
	// 注册中间件
	route.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

func setUp404Handler(route *gin.Engine) {
	// 注册404处理
	route.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "404页面")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    http.StatusNotFound,
				"error_message": "找不到url",
			})
		}
	})
}
