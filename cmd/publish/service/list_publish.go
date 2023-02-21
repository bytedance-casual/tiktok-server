package service

import (
	"context"
	"tiktok-server/cmd/publish/dal/db"
	"tiktok-server/cmd/publish/rpc"
	"tiktok-server/kitex_gen/favorite"
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

	videoIdList := make([]int64, len(videos))
	for i, video := range videos {
		videoIdList[i] = int64(video.ID)
	}

	resp, err := rpc.MCountFavorite(s.ctx, &favorite.MCountVideoFavoriteRequest{VideoIdList: videoIdList})
	if err != nil {
		return nil, err
	}

	countList := resp.CountList
	pVideos := make([]*feed.Video, len(videos))
	for i, video := range videos {
		pVideos[i] = &feed.Video{
			Id:            videoIdList[i],
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: countList[i],
			Title:         video.Title,
		}
	}
	return pVideos, nil
}
