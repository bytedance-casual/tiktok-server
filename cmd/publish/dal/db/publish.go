package db

import (
	"context"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	Title         string
}

func (v *Video) TableName() string {
	return "videos"
}

func CreateVideo(video *Video, ctx context.Context) (int64, error) {
	result := DB.WithContext(ctx).Create(video)
	return int64(video.ID), result.Error
}

func QueryVideo(userId int64, ctx context.Context) ([]*Video, error) {
	resp := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author_id = ?", userId).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func MGetVideos(videoIdList []int64, ctx context.Context) ([]*Video, error) {
	resp := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("id in ?", videoIdList).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}
