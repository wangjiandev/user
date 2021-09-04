package handler

import (
	context "context"
	"github.com/wangjiandev/user/domain/model"
	"github.com/wangjiandev/user/domain/service"
	pb "github.com/wangjiandev/user/proto/user"
)

type User struct {
	UserService service.IUserService
}

func (u *User) Register(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	user := &model.User{
		UserName:     request.UserName,
		FirstName:    request.FirstName,
		HashPassword: request.Pwd,
	}
	_, err := u.UserService.AddUser(user)
	if err != nil {
		return err
	}
	response.Message = "注册成功"
	return nil
}

func (u *User) Login(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	isOk, err := u.UserService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, request *pb.UserInfoRequest, response *pb.UserInfoResponse) error {
	user, err := u.UserService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response = UserForResponse(user)
	return nil
}

func UserForResponse(user *model.User) *pb.UserInfoResponse {
	response := &pb.UserInfoResponse{}
	response.UserId = user.ID
	response.UserName = user.UserName
	response.FirstName = user.FirstName
	return response
}
