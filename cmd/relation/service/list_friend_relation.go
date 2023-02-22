package service

import (
	"context"
	"strconv"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/kitex_gen/relation"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s FriendListService) ListFriend(userId string) ([]*relation.FriendUser, error) {
	var list []db.Follow
	if err := db.DB.WithContext(s.ctx).Model(&db.Follow{}).Where(map[string]any{"user_id": userId, "is_follow": true}).Find(list).Error; err != nil {
		return nil, err
	}
	res := make([]*relation.FriendUser, 0)
	for _, follow := range list {
		r, err := s.hasRelation(follow.FollowedUserId, userId)
		if err != nil {
			return nil, err
		}
		if r {
			friendId, err2 := strconv.ParseInt(follow.FollowedUserId, 10, 64)
			if err2 != nil {
				return nil, err2
			}
			res = append(res, &relation.FriendUser{Id: friendId})
		}
	}
	return res, nil
}

func (s FriendListService) hasRelation(userId string, followedId string) (bool, error) {
	var count int64
	if err := db.DB.WithContext(s.ctx).Model(&db.Follow{}).Where(map[string]any{"user_id": userId, "is_follow": true, "followed_user_id": followedId}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 1, nil
}
