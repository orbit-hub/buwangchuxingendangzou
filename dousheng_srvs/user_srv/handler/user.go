package handler

import (
	"bwcxgdz/v2/user_srv/global"
	"bwcxgdz/v2/user_srv/model"
	"bwcxgdz/v2/user_srv/proto"
	"context"
	"crypto/sha512"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	*proto.UnimplementedUserServer
}

func ModelToRsponse(user model.User) proto.UserInfoResponse {
	userInfoRsp := proto.UserInfoResponse{
		Id:            uint32(user.ID),
		Name:          user.Name,
		Password:      user.Password,
		NickName:      user.NickName,
		FollowerCount: user.FollowerCount,
		FollowCount:   user.FollowCount,
		IsFollow:      true,
	}
	return userInfoRsp
}

func (s *UserServer) GetUserInfo(c context.Context, req *proto.IdRequest) (*proto.UserListResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	rsp := &proto.UserListResponse{}
	userInfoRsp := ModelToRsponse(user)
	rsp.Data = append(rsp.Data, &userInfoRsp)
	return rsp, nil
}

func (s *UserServer) GetUserByName(c context.Context, req *proto.NameRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where("name = ?", req.Name).First(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) CreateUser(c context.Context, req *proto.CreateUserInfoRequest) (*proto.UserInfoResponse, error) {
	//新建用户
	var user model.User
	result := global.DB.Where(&model.User{Name: req.Name}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Name = req.Name
	user.NickName = req.NickName

	//密码加密
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) CheckPassWord(c context.Context, req *proto.PasswordCheckInfoRequest) (*proto.CheckResponse, error) {
	//校验密码
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &proto.CheckResponse{Success: check}, nil
}
