package service

import (
	"context"
	"tiktok-server/cmd/comment/dal/db"
	"tiktok-server/cmd/comment/rpc"
	"tiktok-server/kitex_gen/comment"
	"tiktok-server/kitex_gen/user"
)

type QueryCommentService struct {
	ctx context.Context
}

func NewQueryCommentService(ctx context.Context) *QueryCommentService {
	return &QueryCommentService{
		ctx: ctx,
	}
}

func (s *QueryCommentService) QueryComment(req *comment.CommentListRequest, userId int64) ([]*comment.Comment, error) {
	commentList := make([]*comment.Comment, 0)
	comments, err := db.QueryComment(req.VideoId, s.ctx)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return commentList, nil
	}

	userIdList := make([]int64, len(comments))
	for i, comm := range comments {
		userIdList[i] = comm.UserId
	}

	resp, err := rpc.MGetUsers(s.ctx, &user.UsersMGetRequest{UserId: userId, UserIdList: userIdList})
	if err != nil {
		return nil, err
	}

	users := resp.Users
	for i, comm := range comments {
		commentList = append(commentList, &comment.Comment{
			Id:         int64(comm.ID),
			User:       users[i],
			Content:    comm.Content.Content,
			CreateDate: comm.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return commentList, nil
}
