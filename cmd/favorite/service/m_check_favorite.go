package service

import (
	"context"
	"tiktok-server/cmd/favorite/dal/db"
	"tiktok-server/kitex_gen/favorite"
)

type MCheckFavoriteService struct {
	ctx context.Context
}

func NewMCheckFavoriteService(ctx context.Context) *MCheckFavoriteService {
	return &MCheckFavoriteService{
		ctx: ctx,
	}
}

func (s *MCheckFavoriteService) MCheck(req *favorite.MCheckFavoriteRequest) ([]bool, error) {
	boolList, err := db.MCheckFavorite(req.UserId, req.VideoIdList, s.ctx)
	if err != nil || len(req.VideoIdList) != len(boolList) {
		return nil, err
	}
	return boolList, nil
}
