package service

import (
	"context"
	"tiktok-server/cmd/favorite/dal/db"
	"tiktok-server/kitex_gen/favorite"
	"tiktok-server/kitex_gen/feed"
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
func (s *FavoriteListService) GetFavoriteList(req *favorite.FavoriteListRequest) ([]*feed.Video, error) {
	videos, err := db.ListFavorite(req.UserId, s.ctx)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
