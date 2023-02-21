package main

import (
	"context"
	"tiktok-server/cmd/comment/rpc"
	"tiktok-server/cmd/comment/service"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// ActionComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ActionComment(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.VideoId <= 0 {
		resp = &comment.CommentActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	if req.ActionType == 1 {
		resp, err = doActionComment1(ctx, req)
	} else if req.ActionType == 2 {
		resp, err = doActionComment2(ctx, req)
	} else {
		err = erren.TypeNotSupportErr
	}

	if err != nil {
		errStr := err.Error()
		resp = &comment.CommentActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	return resp, nil
}

// ListComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ListComment(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.VideoId <= 0 {
		resp = &comment.CommentListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	var userId int64
	if len(req.Token) != 0 {
		claims, err := middleware.ParseToken(req.Token)
		if err != nil {
			return nil, err
		}
		userId = claims.ID
	}

	comments, err := service.NewQueryCommentService(ctx).QueryComment(req, userId)
	if err != nil {
		errStr := err.Error()
		resp = &comment.CommentListResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &comment.CommentListResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, CommentList: comments}
	return resp, nil
}

func doActionComment1(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	if len(*req.CommentText) == 0 {
		resp = &comment.CommentActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	user, err := rpc.GetUserFromToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	// TODO 整体事务
	dbComment, err := service.NewActionCommentService(ctx).CreateComment(req, user.Id)
	if err != nil {
		return nil, err
	}

	resp = &comment.CommentActionResponse{
		StatusCode: erren.SuccessCode,
		StatusMsg:  &erren.Success.ErrMsg,
		Comment: &comment.Comment{
			Id:         int64(dbComment.ID),
			User:       user,
			Content:    *req.CommentText,
			CreateDate: dbComment.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}
	return resp, nil
}

func doActionComment2(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	if *req.CommentId <= 0 {
		resp = &comment.CommentActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	// TODO 整体事务
	err = service.NewActionCommentService(ctx).DeleteComment(req)
	if err != nil {
		return nil, err
	}

	resp = &comment.CommentActionResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg}
	return resp, nil
}

// MCountVideoComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MCountVideoComment(ctx context.Context, req *comment.MCountVideoCommentRequest) (resp *comment.MCountVideoCommentResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if len(req.VideoIdList) == 0 {
		resp = &comment.MCountVideoCommentResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	countList, err := service.NewMCountVideoCommentService(ctx).MCount(req)
	if err != nil {
		errStr := err.Error()
		resp = &comment.MCountVideoCommentResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &comment.MCountVideoCommentResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, CountList: countList}
	return resp, nil
}
