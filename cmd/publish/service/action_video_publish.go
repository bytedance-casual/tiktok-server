package service

import (
	"context"
	"tiktok-server/cmd/publish/dal/db"
	"tiktok-server/kitex_gen/publish"
)

type ActionVideoService struct {
	ctx context.Context
}

func NewActionVideoService(ctx context.Context) *ActionVideoService {
	return &ActionVideoService{
		ctx: ctx,
	}
}

func (s *ActionVideoService) UpdateFavorite(req *publish.PublishVideoActionRequest) error {
	err := db.UpdateVideoFavorite(req.VideoId, req.Increase, s.ctx)
	return err
}

func (s *ActionVideoService) UpdateComment(req *publish.PublishVideoActionRequest) error {
	err := db.UpdateVideoComment(req.VideoId, req.Increase, s.ctx)
	return err
}
