package db

import (
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoId int64
	UserId  int64
	Content *Content // 自动注册 belongs-to `gorm:"foreignKey:CommentId"`
}

func (c *Comment) TableName() string {
	return "comments"
}

type Content struct {
	gorm.Model
	CommentId int64
	Content   string
}

func (c *Content) TableName() string {
	return "contents"
}

func CreateComment(comment *Comment, ctx context.Context) (int64, error) {
	result := DB.WithContext(ctx).Create(comment)
	return int64(comment.ID), result.Error
}

func DeleteComment(videoId int64, commentId int64, ctx context.Context) error {
	session := DB.WithContext(ctx)
	session.Model(&Comment{}).Association("Content")
	result := session.Where("video_id = ?", videoId).Select("Content").Delete(&Comment{
		Model: gorm.Model{
			ID: uint(commentId),
		},
	})
	return result.Error
}

func QueryComment(videoId int64, ctx context.Context) ([]*Comment, error) {
	resp := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Preload("Content").Where("video_id = ?", videoId).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}
