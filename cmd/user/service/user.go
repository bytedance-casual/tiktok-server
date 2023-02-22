package service

import (
	"context"
	"tiktok-server/cmd/user/dal/db"
	"tiktok-server/cmd/user/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/relation"
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
	targetUserIdList := []int64{req.UserId}
	users, err := db.MGetUsers(s.ctx, targetUserIdList)
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
	userId := claims.ID

	resp1, err := rpc.MCountRelation(s.ctx, &relation.MCountRelationRequest{UserIdList: targetUserIdList})
	if err != nil {
		return nil, err
	}
	resp2, err := rpc.MCheckFollowRelation(s.ctx, &relation.MCheckFollowRelationRequest{UserId: userId, UserIdList: targetUserIdList})
	if err != nil {
		return nil, err
	}
	followCount := resp1.FollowCountList[0]
	followerCount := resp1.FollowerCountList[0]
	isFollow := resp2.CheckList[0]

	u := user.User{
		Id:            int64(users[0].ID),
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
		Name:          users[0].Username,
	}
	return &u, nil
}
