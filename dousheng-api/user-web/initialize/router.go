package initialize

import (
	"bwcxgdz/api/user-web/middlewares"
	"bwcxgdz/api/user-web/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors()) //跨域
	ApiGroup := Router.Group("/douyin")
	router.InitUserRouter(ApiGroup)
	return Router
}
