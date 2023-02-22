package service

import (
	"context"
	"tiktok-server/cmd/message/dal/db"
	"tiktok-server/kitex_gen/message"
)

type MGetLatestMessageService struct {
	ctx context.Context
}

func NewMGetLatestMessageService(ctx context.Context) *MGetLatestMessageService {
	return &MGetLatestMessageService{
		ctx: ctx,
	}
}

func (s *MGetLatestMessageService) MGetLatestMessage(req *message.MGetLatestMessageRequest) ([]bool, []string, error) {
	userId := req.UserId
	typeList := make([]bool, len(req.FriendIdList))
	contentList := make([]string, len(req.FriendIdList))
	for i, friendId := range req.FriendIdList {
		latestMessage, err := db.QueryLatestMessage(s.ctx, userId, friendId)
		if err != nil {
			return nil, nil, err
		}
		typeList[i] = latestMessage.FromUserId == userId
		contentList[i] = latestMessage.Content
	}
	return typeList, contentList, nil
}
