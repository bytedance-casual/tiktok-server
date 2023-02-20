package service

import (
	"context"
	"tiktok-server/cmd/favorite/dal/db"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/favorite"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	// 1-点赞
	if req.ActionType == 1 {
		return db.AddFavoriteAction(s.ctx, req.UserId, req.VideoId)
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		return db.CancelFavoriteAction(s.ctx, req.UserId, req.VideoId)
	}
	return erren.ServiceErr

}
