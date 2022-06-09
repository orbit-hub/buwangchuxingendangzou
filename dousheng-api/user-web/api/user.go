package api

import (
	"bwcxgdz/api/user-web/global"
	"bwcxgdz/api/user-web/global/response"
	"bwcxgdz/api/user-web/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"status_msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"status_msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_msg": e.Code(),
				})
			}
			return
		}
	}
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func GetUserInfo(ctx *gin.Context) {

	//拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d",
		global.ServerConfig.UserSrvInfo.Host,
		global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserInfo] 连接用户服务失败", "msg", err.Error())
	}
	//调用接口
	userSrvClient := proto.NewUserClient(userConn)

	userId := ctx.DefaultQuery("user_id", "")
	UserIdInt, _ := strconv.Atoi(userId)
	//token := ctx.DefaultQuery("token", "")

	rsp, err := userSrvClient.GetUserInfo(context.Background(), &proto.IdRequest{
		Id: uint32(UserIdInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserInfo] 查询 【用户列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	value := rsp.Data[0]

	user := response.User{
		Id:            value.Id,
		Name:          value.Name,
		FollowerCount: value.FollowerCount,
		FollowCount:   value.FollowCount,
		IsFollow:      value.IsFollow,
	}
	userResponse := response.UserResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		User: user,
	}

	ctx.JSON(http.StatusOK, &userResponse)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d",
		global.ServerConfig.UserSrvInfo.Host,
		global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserInfo] 连接用户服务失败", "msg", err.Error())
	}
	//调用接口
	userSrvClient := proto.NewUserClient(userConn)

	userSrvClient.GetUserInfo()

}
