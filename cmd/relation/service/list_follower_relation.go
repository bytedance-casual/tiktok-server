package service

import (
	"context"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/relation"
	"tiktok-server/kitex_gen/user"
)

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{
		ctx: ctx,
	}
}

// 查询粉丝列表
func (s *FollowerListService) GetFollowerList(req *relation.RelationFollowerListRequest) ([]*user.User, error) {
	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	users, err := db.QueryFollowerList(s.ctx, claims.ID)
	if err != nil {
		return nil, err
	}
	myfollowlist, err := db.QueryFollowList(s.ctx, claims.ID)
	if err != nil {
		return nil, err
	}
	//获取我的粉丝列表和关注列表，遍历查询我是否关注了某粉丝
	for _, v := range users { //v是结构体数组中的一个结构体实例，我的粉丝
		v.IsFollow = false
		for _, j := range myfollowlist { //j是结构体数组中的一个结构体实例，我的关注
			if v.Id == j.Id {
				v.IsFollow = true
			}
		}
	}
	//users2 := make([]*user.User, 0)
	//users2 = append(users2, &user.User{Name: "2"})
	//u := user.User{Id: int64(users[0].ID), FollowCount: users[0].FollowCount, FollowerCount: users[0].FollowerCount, IsFollow: isFollow, Name: users[0].Username}
	//u := users[0]
	return users, nil
}
