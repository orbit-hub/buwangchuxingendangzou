package main

import (
	"bwcxgdz/api/user-web/global"
	"bwcxgdz/api/user-web/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {

	//初始化；logger
	initialize.InitLogger()
	//初始化配置文件
	initialize.InitConfig()
	//初始化routers
	Router := initialize.Routers()
	//初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	zap.S().Infof("启动服务器，端口: %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}

}
