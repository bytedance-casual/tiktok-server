package service

import (
	"context"
	"tiktok-server/cmd/favorite/dal/db"
	"tiktok-server/kitex_gen/favorite"
	"tiktok-server/kitex_gen/feed"
	"tiktok-server/kitex_gen/user"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService creates a new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// GetFavoriteList returns a Favorite List
func (s *FavoriteListService) GetFavoriteList(req *favorite.FavoriteListRequest, user *user.User) ([]*feed.Video, error) {
	videos, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	fVideos := make([]*feed.Video, len(videos))
	for i, video := range videos {
		fVideos[i] = &feed.Video{
			Id:            int64(video.Id),
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		}
	}
	return fVideos, nil
}
