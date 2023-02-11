package db

import (
	"context"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromUserId int64
	ToUserId   int64
	Content    string
}

func (m *Message) TableName() string {
	return "messages"
}

// CreateMessage create a message record
func CreateMessage(ctx context.Context, message *Message) (int64, error) {
	result := DB.WithContext(ctx).Create(message)
	return int64(message.ID), result.Error
}

// QueryMessage with specific form & to user_id
func QueryMessage(ctx context.Context, fromUserId int64, toUserId int64) ([]*Message, error) {
	resp := make([]*Message, 0)
	if err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}
