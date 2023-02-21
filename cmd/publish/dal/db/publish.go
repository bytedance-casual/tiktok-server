package db

import (
	"context"
	"gorm.io/gorm"
	"tiktok-server/internal/utils"
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

// UpdateVideoFavorite update video favorite number by (increase ? +1 : -1)
func UpdateVideoFavorite(videoId int64, increase bool, ctx context.Context) error {
	opt := utils.Ternary(increase, "favorite_count + ?", "favorite_count - ?")
	result := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr(opt, 1))
	return result.Error
}

// UpdateVideoComment update video comment number by (increase ? +1 : -1)
func UpdateVideoComment(videoId int64, increase bool, ctx context.Context) error {
	opt := utils.Ternary(increase, "comment_count + ?", "comment_count - ?")
	result := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr(opt, 1))
	return result.Error
}
