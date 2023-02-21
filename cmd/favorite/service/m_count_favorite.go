package service

import (
	"context"
	"tiktok-server/cmd/favorite/dal/db"
	"tiktok-server/kitex_gen/favorite"
)

type MCountFavoriteService struct {
	ctx context.Context
}

func NewMCountFavoriteService(ctx context.Context) *MCountFavoriteService {
	return &MCountFavoriteService{
		ctx: ctx,
	}
}

func (s *MCountFavoriteService) MCount(req *favorite.MCountVideoFavoriteRequest) ([]int64, error) {
	countList, err := db.MCountFavorite(req.VideoIdList, s.ctx)
	if err != nil || len(req.VideoIdList) != len(countList) {
		return nil, err
	}
	return countList, nil
}
