package service

import (
	"context"
	"tiktok-server/cmd/comment/dal/db"
	"tiktok-server/kitex_gen/comment"
)

type MCountVideoCommentService struct {
	ctx context.Context
}

func NewMCountVideoCommentService(ctx context.Context) *MCountVideoCommentService {
	return &MCountVideoCommentService{
		ctx: ctx,
	}
}

func (s *MCountVideoCommentService) MCount(req *comment.MCountVideoCommentRequest) ([]int64, error) {
	countList, err := db.MCountComment(req.VideoIdList, s.ctx)
	if err != nil || len(countList) != len(req.VideoIdList) {
		return nil, err
	}
	return countList, nil
}
