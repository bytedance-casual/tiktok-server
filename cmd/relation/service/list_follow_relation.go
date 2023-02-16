package service

import (
	"context"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/relation"
	"tiktok-server/kitex_gen/user"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{
		ctx: ctx,
	}
}

// 查询关注列表
func (s *FollowListService) GetFollowList(req *relation.RelationFollowListRequest) ([]*user.User, error) {
	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	users, err := db.QueryFollowList(s.ctx, claims.ID)

	if err != nil {
		return nil, err
	}

	//users2 := make([]*user.User, 0)
	//users2 = append(users2, &user.User{Name: "1"})
	//u := user.User{Id: int64(users[0].ID), FollowCount: users[0].FollowCount, FollowerCount: users[0].FollowerCount, IsFollow: isFollow, Name: users[0].Username}
	//u := users[0]
	return users, nil
}
