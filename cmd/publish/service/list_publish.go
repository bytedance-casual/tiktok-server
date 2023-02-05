package service

import (
	"context"
	"tiktok-server/cmd/publish/dal/db"
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

func (s *QueryVideoService) QueryVideo(req *publish.PublishListRequest, user *user.User) ([]*publish.Video, error) {
	videos, err := db.QueryVideo(req.UserId, s.ctx)
	if err != nil {
		return nil, err
	}
	// TODO /favorite/list/
	pUser := publish.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}

	pVideos := make([]*publish.Video, len(videos))
	for i, video := range videos {
		pVideos[i] = &publish.Video{
			Id:            int64(video.ID),
			Author:        &pUser,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			// TODO
			IsFavorite: false,
			Title:      video.Title,
		}
	}
	return pVideos, nil
}
