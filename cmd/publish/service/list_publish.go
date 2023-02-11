package service

import (
	"context"
	"tiktok-server/cmd/publish/dal/db"
	"tiktok-server/kitex_gen/feed"
	"tiktok-server/kitex_gen/publish"
	"tiktok-server/kitex_gen/user"
)

type QueryVideoService struct {
	ctx context.Context
}

func NewQueryVideoService(ctx context.Context) *QueryVideoService {
	return &QueryVideoService{
		ctx: ctx,
	}
}

func (s *QueryVideoService) QueryVideo(req *publish.PublishListRequest, user *user.User) ([]*feed.Video, error) {
	videos, err := db.QueryVideo(req.UserId, s.ctx)
	if err != nil {
		return nil, err
	}

	pVideos := make([]*feed.Video, len(videos))
	for i, video := range videos {
		pVideos[i] = &feed.Video{
			Id:            int64(video.ID),
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			// TODO /favorite/list/
			IsFavorite: false,
			Title:      video.Title,
		}
	}
	return pVideos, nil
}
