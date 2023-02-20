package db

import (
	"context"
	"gorm.io/gorm"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/feed"
	user2 "tiktok-server/kitex_gen/user"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (f *Favorite) TableName() string {
	return "favorite"
}

// GetFavoriteRelation 获取用户视频
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (*feed.Video, error) {
	user := new(user2.User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	video := new(feed.Video)
	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&video, vid); err != nil {
		return nil, err
	}
	return video, nil
}

// AddFavoriteAction 点赞操作
func AddFavoriteAction(ctx context.Context, userId int64, videoId int64) error {
	favorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	//用户视频点赞数
	var cnt int64 = 0
	err := DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&cnt).Error
	if err != nil {
		return err
	}
	if cnt != 0 {
		return nil
	}

	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&favorite).Error; err != nil {
			return err
		}
		if err := DB.WithContext(ctx).Model(&feed.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// CancelFavoriteAction 取消点赞操作
func CancelFavoriteAction(ctx context.Context, userId int64, videoId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//删除点赞数据
		user := new(user2.User)
		if err := tx.WithContext(ctx).First(user, userId).Error; err != nil {
			return err
		}

		video, err := GetFavoriteRelation(ctx, userId, videoId)
		if err != nil {
			return err
		}

		err = tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideos").Delete(video)
		if err != nil {
			return err
		}

		//改变 video 表中的 favorite count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return erren.ParamErr
		}

		return nil
	})
	return err
}

// FavoriteList 返回userid用户点赞的视频列表
func FavoriteList(ctx context.Context, uid int64) ([]*feed.Video, error) {
	user := new(user2.User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	//videos := []feed.Video{}
	videos := make([]*feed.Video, 0)
	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&videos); err != nil {
		return nil, err
	}
	return videos, nil
}
