package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"tiktok-server/cmd/user/dal/db"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/user"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.UserRegisterRequest) (int64, string, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	var token string
	if err != nil {
		return 0, "", err
	}
	if len(users) != 0 {
		return 0, "", erren.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, "", err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	userid, err := db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: password,
	}})
	if err != nil {
		return 0, "", err
	}
	if token, err = middleware.GenToken(req.Username, int64(userid)); err != nil {
		return 0, "", err
	}
	return int64(userid), token, nil
}
