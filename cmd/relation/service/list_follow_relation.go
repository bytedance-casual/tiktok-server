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

	newusers := make([]*user.User, 0) //不同结构体，解耦，重新赋值返回

	for k, v := range users {
		newusers[k].Id = v.Id
		newusers[k].Name = v.Name
		newusers[k].FollowerCount = v.FollowerCount
		newusers[k].FollowCount = v.FollowCount
		newusers[k].IsFollow = v.IsFollow
	}
	//users2 = append(users2, &user.User{Name: "1"})
	//u := user.User{Id: int64(users[0].ID), FollowCount: users[0].FollowCount, FollowerCount: users[0].FollowerCount, IsFollow: isFollow, Name: users[0].Username}
	//u := users[0]
	return newusers, nil
}
