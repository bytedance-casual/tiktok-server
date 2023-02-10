package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

func (table *Video) TableName() string {
	return "videos"
}

// MGetVideosByTime 根据时间返回视频
func MGetVideosByTime(ctx context.Context, authorId int64, latestTime *int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	if latestTime == nil || *latestTime == 0 {
		curTime := time.Now().UnixMilli()
		latestTime = &curTime
	}

	videoTime := time.UnixMilli(*latestTime).Format("2006-01-02 15:04:05")
	if err := DB.WithContext(ctx).Where("author_id != ? AND created_at < ?", authorId, videoTime).Order("created_at desc").Limit(30).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
