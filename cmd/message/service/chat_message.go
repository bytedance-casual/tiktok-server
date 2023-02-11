package service

import (
	"context"
	"tiktok-server/cmd/message/dal/db"
	"tiktok-server/kitex_gen/message"
)

type GetMessageService struct {
	ctx context.Context
}

func NewGetMessageService(ctx context.Context) *GetMessageService {
	return &GetMessageService{
		ctx: ctx,
	}
}

func (s *GetMessageService) GetMessage(req *message.MessageChatRequest, userId int64) ([]*message.Message, error) {
	dbMessages, err := db.QueryMessage(s.ctx, userId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	messages := make([]*message.Message, len(dbMessages))
	for i, dbMessage := range dbMessages {
		timeFormat := dbMessage.CreatedAt.Format("2006-01-02 15:04:05")
		messages[i] = &message.Message{
			Id:         int64(dbMessage.ID),
			ToUserId:   dbMessage.ToUserId,
			FromUserId: dbMessage.FromUserId,
			Content:    dbMessage.Content,
			CreateTime: &timeFormat,
		}
	}
	return messages, nil
}
