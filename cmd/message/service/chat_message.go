package service

import (
	"context"
	"sort"
	"strconv"
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
	this2ThatMessages, err := db.QueryMessage(s.ctx, userId, req.ToUserId, req.PreMsgTime)
	if err != nil {
		return nil, err
	}
	that2ThisMessages, err := db.QueryMessage(s.ctx, req.ToUserId, userId, req.PreMsgTime)
	if err != nil {
		return nil, err
	}

	dbMessages := append(this2ThatMessages, that2ThisMessages...)
	sort.SliceStable(dbMessages, func(i, j int) bool { // 升序排序
		return dbMessages[i].CreatedAt.UnixMilli() < dbMessages[j].CreatedAt.UnixMilli()
	})

	messages := make([]*message.Message, len(dbMessages))
	for i, dbMessage := range dbMessages {
		timeFormat := strconv.FormatInt(dbMessage.CreatedAt.UnixMilli(), 10)
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
