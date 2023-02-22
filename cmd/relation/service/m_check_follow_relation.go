package service

import (
	"context"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/kitex_gen/relation"
)

type MCheckFollowRelationService struct {
	ctx context.Context
}

func NewMCheckFollowRelationService(ctx context.Context) *MCheckFollowRelationService {
	return &MCheckFollowRelationService{
		ctx: ctx,
	}
}

func (s *MCheckFollowRelationService) MCheckFollow(req *relation.MCheckFollowRelationRequest) ([]bool, error) {
	checkList, err := db.MCheckFollow(req.UserId, req.UserIdList, s.ctx)
	if err != nil || len(req.UserIdList) != len(checkList) {
		return nil, err
	}
	return checkList, nil
}
