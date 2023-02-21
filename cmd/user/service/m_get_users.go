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
func (s *MGetUsersService) MGetUsers(req *user.UsersMGetRequest) (map[int64]*user.User, error) {
	users, err := db.MGetUsers(s.ctx, req.UserIdList)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, erren.UserNotExistErr
	}

	userMap := make(map[int64]*user.User, len(users))
	for _, dbUser := range users {
		// TODO move to rpc.IsFollow
		//isFollow, err := db.QueryIsFollow(s.ctx, req.UserId, int64(dbUser.ID))
		//if err != nil {
		//	return nil, err
		//}
		id := int64(dbUser.ID)
		userMap[id] = &user.User{
			Id:   id,
			Name: dbUser.Username,
			// TODO
			//FollowCount:   dbUser.FollowCount,
			//FollowerCount: dbUser.FollowerCount,
			IsFollow: false,
		}
	}
	return userMap, nil
}
