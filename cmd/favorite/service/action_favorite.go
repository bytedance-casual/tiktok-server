package service

import (
	"context"
	"tiktok-server/cmd/favorite/dal/db"
	"tiktok-server/kitex_gen/favorite"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) AddFavorite(req *favorite.FavoriteActionRequest, userId int64) error {
	_, err := db.AddFavorite(&db.Like{
		UserId:  userId,
		VideoId: req.VideoId,
	}, s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavoriteActionService) CancelFavorite(req *favorite.FavoriteActionRequest, userId int64) error {
	err := db.CancelFavorite(&db.Like{
		UserId:  userId,
		VideoId: req.VideoId,
	}, s.ctx)
	if err != nil {
		return err
	}
	return nil
}
