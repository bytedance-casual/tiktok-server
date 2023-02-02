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

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *user.UserLoginRequest) (int64, string, error) {
	h := md5.New()
	var Token string
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, "", err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, "", err
	}
	if len(users) == 0 {
		return 0, "", erren.AuthorizationFailedErr
	}
	if users[0].Password != passWord {
		return 0, "", erren.AuthorizationFailedErr
	}
	if Token, err = middleware.GenToken(users[0].Username, int64(users[0].ID)); err != nil {
		return 0, "", err
	}
	return int64(users[0].ID), Token, nil
}
