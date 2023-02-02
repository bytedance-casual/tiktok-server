package service

import (
	"context"
	"tiktok-server/cmd/user/dal/db"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/user"
)

type UserService struct {
	ctx context.Context
}

// NewUserService new UserService
func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

// User  user info
func (s *UserService) User(req *user.UserRequest) (*user.User, error) {
	users, err := db.MGetUsers(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, erren.UserNotExistErr
	}
	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	isFollow, err := db.QueryIsFollow(s.ctx, claims.ID, req.UserId)
	if err != nil {
		return nil, err
	}
	u := user.User{Id: int64(users[0].ID), FollowCount: users[0].FollowCount, FollowerCount: users[0].FollowerCount, IsFollow: isFollow, Name: users[0].Username}
	//u := users[0]
	return &u, nil
}
