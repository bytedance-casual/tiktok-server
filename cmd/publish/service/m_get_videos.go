package service

import (
	"context"
	"tiktok-server/cmd/publish/dal/db"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/feed"
	"tiktok-server/kitex_gen/publish"
	"tiktok-server/kitex_gen/user"
)

type MGetVideosService struct {
	ctx context.Context
}

// NewMGetVideosService new MGetVideosService
func NewMGetVideosService(ctx context.Context) *MGetVideosService {
	return &MGetVideosService{
		ctx: ctx,
	}
}

func (s *MGetVideosService) MGetVideos(req *publish.VideosMGetRequest) ([]*feed.Video, error) {
	videos, err := db.MGetVideos(req.VideoIdList, s.ctx)
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 || len(videos) != len(req.VideoIdList) {
		return nil, erren.VideoNotFoundErr
	}

	// 此处视频信息作为一个基础，上层rpc调用时如需其他数据再发起即可
	pVideos := make([]*feed.Video, len(videos))
	for i, video := range videos {
		pVideos[i] = &feed.Video{
			Id:       int64(video.ID),
			Author:   &user.User{},
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			//FavoriteCount: video.FavoriteCount,
			//CommentCount: video.CommentCount,
			//IsFavorite:   false,
			Title: video.Title,
		}
	}
	return pVideos, nil
}
