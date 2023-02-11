package service

import (
	"context"
	"gorm.io/gorm"
	"tiktok-server/cmd/message/dal/db"
	"tiktok-server/kitex_gen/message"
)

type CreateMessageService struct {
	ctx context.Context
}

func NewCreateMessageService(ctx context.Context) *CreateMessageService {
	return &CreateMessageService{
		ctx: ctx,
	}
}

func (s *CreateMessageService) CreateMessage(req *message.MessageActionRequest, userId int64) error {
	_, err := db.CreateMessage(s.ctx, &db.Message{
		Model:      gorm.Model{},
		FromUserId: userId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	})
	if err != nil {
		return err
	}
	return nil
}
