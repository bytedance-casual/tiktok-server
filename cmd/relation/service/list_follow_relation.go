package service

import (
	"context"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/cmd/relation/rpc"
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

// GetFollowList 查询关注列表
func (s *FollowListService) GetFollowList(req *relation.RelationFollowListRequest) ([]*user.User, error) {
	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	userId := claims.ID
	userIdList, err := db.QueryFollowList(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	var userMap map[int64]*user.User
	if len(userIdList) != 0 {
		resp, err := rpc.MGetUsers(s.ctx, &user.UsersMGetRequest{UserId: userId, UserIdList: userIdList})
		if err != nil {
			return nil, err
		}
		userMap = resp.Users
	}

	i := 0
	users := make([]*user.User, len(userMap))
	for _, v := range userMap {
		users[i] = v
		i ++
	}
	return users, nil
}
