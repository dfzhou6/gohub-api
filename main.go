package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	bstConfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	bstConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	bootstrap.SetupLogger()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()
	bootstrap.SetupRoute(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
