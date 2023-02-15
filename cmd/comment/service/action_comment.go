package service

import (
	"context"
	"tiktok-server/cmd/comment/dal/db"
	"tiktok-server/kitex_gen/comment"
)

type ActionCommentService struct {
	ctx context.Context
}

func NewActionCommentService(ctx context.Context) *ActionCommentService {
	return &ActionCommentService{
		ctx: ctx,
	}
}

func (s *ActionCommentService) CreateComment(req *comment.CommentActionRequest, userId int64) (*db.Comment, error) {
	comm := db.Comment{
		VideoId: req.VideoId,
		UserId:  userId,
		Content: &db.Content{
			Content: *req.CommentText,
		},
	}
	_, err := db.CreateComment(&comm, s.ctx)
	if err != nil {
		return nil, err
	}
	return &comm, nil
}

func (s *ActionCommentService) DeleteComment(req *comment.CommentActionRequest) error {
	err := db.DeleteComment(req.VideoId, *req.CommentId, s.ctx)
	if err != nil {
		return err
	}
	return nil
}
