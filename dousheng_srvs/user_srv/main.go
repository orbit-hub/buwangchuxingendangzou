package main

import (
	"bwcxgdz/v2/user_srv/global"
	"bwcxgdz/v2/user_srv/handler"
	"bwcxgdz/v2/user_srv/initialize"
	"bwcxgdz/v2/user_srv/proto"
	"bwcxgdz/v2/user_srv/utils"
	"fmt"
	"net"

	"google.golang.org/grpc/health"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	serverConfig := global.ServerConfig
	IP := serverConfig.Host
	Port := serverConfig.Port
	zap.S().Info("ip: ", IP)
	if Port == 0 {
		Port, _ = utils.GetFreePort()
	}
	zap.S().Info("port: ", Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, Port))

	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
