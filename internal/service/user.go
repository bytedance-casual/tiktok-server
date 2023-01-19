package service

import (
	"errors"
	"tiktok-server/internal/model"
	"tiktok-server/internal/repo"
)

// RegisterUser 注册用户
func RegisterUser(request *model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	if request.Username == "" || request.Password == "" {
		return nil, errors.New("参数不足")
	}
	if repo.QueryUserByName(request.Username) != nil {
		return nil, errors.New("用户已存在")
	}

}
