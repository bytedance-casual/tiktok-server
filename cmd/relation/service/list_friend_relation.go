package service

import (
	"context"
	"strconv"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/cmd/relation/rpc"
	"tiktok-server/internal/utils"
	"tiktok-server/kitex_gen/message"
	"tiktok-server/kitex_gen/relation"
	"tiktok-server/kitex_gen/user"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s FriendListService) ListFriend(userId int64) ([]*relation.FriendUser, error) {
	list := make([]*db.Follow, 0)
	if err := db.DB.WithContext(s.ctx).Model(&db.Follow{}).Where(map[string]any{"user_id": userId, "is_follow": true}).Find(&list).Error; err != nil {
		return nil, err
	}

	friendIdList := make([]int64, 0)
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
			friendIdList = append(friendIdList, friendId)
		}
	}

	resp1, err := rpc.MGetUsers(s.ctx, &user.UsersMGetRequest{UserId: userId, UserIdList: friendIdList})
	if err != nil {
		return nil, err
	}
	resp2, err := rpc.MGetLatestMessage(s.ctx, &message.MGetLatestMessageRequest{UserId: userId, FriendIdList: friendIdList})
	if err != nil {
		return nil, err
	}

	i := 0
	userMap := resp1.Users
	typeList := resp2.TypeList
	contentList := resp2.ContentList
	friendList := make([]*relation.FriendUser, len(userMap))
	for _, u := range userMap {
		friendList[i] = &relation.FriendUser{
			Message:       &contentList[i],
			MsgType:       utils.Ternary[int64](typeList[i], 0, 1),
			Id:            friendIdList[i],
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
		}
		i++
	}
	return friendList, nil
}

func (s FriendListService) hasRelation(userId string, followedId int64) (bool, error) {
	var count int64
	if err := db.DB.WithContext(s.ctx).Model(&db.Follow{}).Where(map[string]any{"user_id": userId, "is_follow": true, "followed_user_id": followedId}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 1, nil
}
