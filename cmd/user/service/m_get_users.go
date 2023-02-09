package service

import (
	"context"
	"tiktok-server/cmd/user/dal/db"
	"tiktok-server/internal/erren"
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
func (s *MGetUsersService) MGetUsers(req *user.UsersMGetRequest) ([]*user.User, error) {
	users, err := db.MGetUsers(s.ctx, req.UserIdList)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 || len(users) != len(req.UserIdList) {
		return nil, erren.UserNotExistErr
	}

	userList := make([]*user.User, 0)
	for i, dbUser := range users {
		// TODO move to rpc.ISFollow
		isFollow, err := db.QueryIsFollow(s.ctx, req.UserId, int64(dbUser.ID))
		if err != nil {
			return nil, err
		}
		userList[i] = &user.User{Id: int64(dbUser.ID), Name: dbUser.Username, FollowCount: dbUser.FollowCount, FollowerCount: dbUser.FollowerCount, IsFollow: isFollow}
	}
	return userList, nil
}
