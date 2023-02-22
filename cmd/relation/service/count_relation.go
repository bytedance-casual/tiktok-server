package service

import (
	"context"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/kitex_gen/relation"
)

type MCountRelationService struct {
	ctx context.Context
}

func NewMCountRelationService(ctx context.Context) *MCountRelationService {
	return &MCountRelationService{
		ctx: ctx,
	}
}

func (s *MCountRelationService) MCount(req *relation.MCountRelationRequest) (followCountList []int64, followerCountList []int64, err error) {
	followCountList, err = db.MCountFollow(req.UserIdList, s.ctx)
	if err != nil {
		return
	}
	followerCountList, err = db.MCountFollower(req.UserIdList, s.ctx)
	if err != nil {
		return
	}
	return
}
