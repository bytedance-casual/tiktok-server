package db

import (
	"context"
	"gorm.io/gorm"
	"time"
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
func QueryMessage(ctx context.Context, fromUserId int64, toUserId int64, preMsgTime int64) ([]*Message, error) {
	resp := make([]*Message, 0)
	millTime := time.UnixMilli(preMsgTime)
	if err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ? AND created_at > ?", fromUserId, toUserId, millTime).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryLatestMessage query latest message from from-to-conversion
func QueryLatestMessage(ctx context.Context, userId1 int64, userId2 int64) (*Message, error) {
	var message *Message
	if err := DB.WithContext(ctx).Where("`created_at` = (SELECT max( `created_at` ) FROM `messages` WHERE (`from_user_id` = ? AND `to_user_id` = ?) OR (`from_user_id` = ? AND `to_user_id` = ?))", userId1, userId2, userId2, userId1).Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}
