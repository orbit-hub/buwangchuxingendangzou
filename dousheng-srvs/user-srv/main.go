package main

import (
	"bwcxgdz/v2/user-srv/handler"
	"bwcxgdz/v2/user-srv/proto"
	"bwcxgdz/v2/user-srv/utils"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口号")
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}