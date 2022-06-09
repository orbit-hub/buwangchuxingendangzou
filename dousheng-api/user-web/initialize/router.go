package initialize

import (
	"bwcxgdz/api/user-web/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/douyin")
	router.InitUserRouter(ApiGroup)
	return Router
}
