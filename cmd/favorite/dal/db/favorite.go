package db

import (
	"context"
	"gorm.io/gorm"
	"tiktok-server/cmd/favorite/rpc"
	"tiktok-server/internal/utils"
	"tiktok-server/kitex_gen/feed"
	"tiktok-server/kitex_gen/publish"
)

type Like struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (f *Like) TableName() string {
	return "likes"
}

// AddFavorite 点赞操作
func AddFavorite(like *Like, ctx context.Context) (int64, error) {
	result := DB.WithContext(ctx).Create(like)
	return int64(like.ID), result.Error
}

// CancelFavorite 取消点赞操作
func CancelFavorite(like *Like, ctx context.Context) error {
	result := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", like.UserId, like.VideoId).Delete(like)
	return result.Error
}

// ListFavorite 返回userid用户点赞的视频列表
func ListFavorite(userId int64, ctx context.Context) ([]*feed.Video, error) {
	resp := make([]*Like, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&resp).Error; err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return make([]*feed.Video, 0), nil
	}

	videoIdList := make([]int64, len(resp))
	for i, like := range resp {
		videoIdList[i] = like.VideoId
	}

	countList, err := MCountFavorite(videoIdList, ctx)
	if err != nil {
		return nil, err
	}

	mresp, err := rpc.MGetVideos(ctx, &publish.VideosMGetRequest{
		UserId:      userId,
		VideoIdList: videoIdList,
	})
	if err != nil {
		return nil, err
	}

	for i, video := range mresp.Videos {
		video.IsFavorite = true
		video.FavoriteCount = countList[i]
	}
	return mresp.Videos, nil
}

func MCountFavorite(videoIdList []int64, ctx context.Context) ([]int64, error) {
	countList := make([]int64, len(videoIdList))
	for i, videoId := range videoIdList {
		if err := DB.WithContext(ctx).Model(&Like{}).Where("video_id = ?", videoId).Count(&countList[i]).Error; err != nil {
			return nil, err
		}
	}
	return countList, nil
}

func MCheckFavorite(userId int64, videoIdList []int64, ctx context.Context) ([]bool, error) {
	set := utils.NewSet[int64]()
	likes := make([]*Like, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id IN ?", userId, videoIdList).Find(&likes).Error; err != nil {
		return nil, err
	}
	for _, like := range likes {
		set.Add(like.VideoId)
	}
	boolList := make([]bool, len(videoIdList))
	for i, videoId := range videoIdList {
		boolList[i] = set.Contains(videoId)
	}
	return boolList, nil
}
