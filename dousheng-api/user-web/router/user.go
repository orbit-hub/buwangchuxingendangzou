package router

import (
	"bwcxgdz/api/user-web/api"
	"bwcxgdz/api/user-web/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		//UserRouter.GET("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("login", api.Login)
		UserRouter.POST("register", api.Register)
		UserRouter.GET("", middlewares.JWTAuth(), api.GetUserInfo)
	}
	//服务注册和发现
}
