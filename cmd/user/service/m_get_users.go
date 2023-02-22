package service

import (
	"context"
	"tiktok-server/cmd/user/dal/db"
	"tiktok-server/cmd/user/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/relation"
	"tiktok-server/kitex_gen/user"
)

type MGetUsersService struct {
	ctx context.Context
}

// NewMGetUsersService new MGetUsersService
func NewMGetUsersService(ctx context.Context) *MGetUsersService {
	return &MGetUsersService{
		ctx: ctx,
	}
}

// MGetUsers batch get users
func (s *MGetUsersService) MGetUsers(req *user.UsersMGetRequest) (map[int64]*user.User, error) {
	users, err := db.MGetUsers(s.ctx, req.UserIdList)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, erren.UserNotExistErr
	}

	userIdList := make([]int64, len(users))
	for i, u := range users {
		userIdList[i] = int64(u.ID)
	}

	resp1, err := rpc.MCountRelation(s.ctx, &relation.MCountRelationRequest{UserIdList: userIdList})
	if err != nil {
		return nil, err
	}

	isFollowList := make([]bool, len(userIdList))
	if req.UserId != 0 {
		resp2, err := rpc.MCheckFollowRelation(s.ctx, &relation.MCheckFollowRelationRequest{UserId: req.UserId, UserIdList: userIdList})
		if err != nil {
			return nil, err
		}
		isFollowList = resp2.CheckList
	}

	followCountList := resp1.FollowCountList
	followerCountList := resp1.FollowerCountList
	userMap := make(map[int64]*user.User, len(users))
	for i, dbUser := range users {
		id := userIdList[i]
		userMap[id] = &user.User{
			Id:            id,
			Name:          dbUser.Username,
			FollowCount:   followCountList[i],
			FollowerCount: followerCountList[i],
			IsFollow:      isFollowList[i],
		}
	}
	return userMap, nil
}
