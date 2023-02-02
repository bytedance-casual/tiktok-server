package main

import (
	"context"
	"tiktok-server/cmd/user/service"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// User implements the UserServiceImpl interface.
func (s *UserServiceImpl) User(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp = new(user.UserResponse)

	if req.UserId <= 0 || len(req.Token) == 0 {
		resp = &user.UserResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	users, err := service.NewUserService(ctx).User(req)
	if err != nil {
		errstr := err.Error()
		resp = &user.UserResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errstr}
		return resp, nil
	}
	resp = &user.UserResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, User: users}
	return resp, nil
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = &user.UserRegisterResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	userId, token, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		errstr := err.Error()
		resp = &user.UserRegisterResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errstr}
		return resp, nil
	}
	resp = &user.UserRegisterResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, UserId: int64(userId), Token: token}
	return resp, nil

}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = &user.UserLoginResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	userId, token, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		errstr := err.Error()
		resp = &user.UserLoginResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errstr}
		return resp, nil
	}
	resp = &user.UserLoginResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, UserId: int64(userId), Token: token}
	return resp, nil
}
