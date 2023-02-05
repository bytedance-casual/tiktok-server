package main

import (
	"context"
	"tiktok-server/cmd/publish/rpc"
	"tiktok-server/cmd/publish/service"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/publish"
	"tiktok-server/kitex_gen/user"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// ActionPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ActionPublish(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if len(req.Token) == 0 || len(req.Data) == 0 || len(req.Title) == 0 {
		resp = &publish.PublishActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, err
	}

	tokenUser, err := getUserFromToken(req.Token, ctx)
	if err != nil {
		errMsg := err.Error()
		resp = &publish.PublishActionResponse{StatusCode: erren.AuthorizationFailedErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	err = service.NewUploadVideoService(ctx).UploadVideo(req, tokenUser)
	if err != nil {
		errMsg := err.Error()
		resp = &publish.PublishActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	resp = &publish.PublishActionResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg}
	return resp, nil
}

// ListPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ListPublish(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if len(req.Token) == 0 || req.UserId <= 0 {
		resp = &publish.PublishListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, err
	}

	tokenUser, err := getUserFromToken(req.Token, ctx)
	if err != nil {
		errMsg := err.Error()
		resp = &publish.PublishListResponse{StatusCode: erren.AuthorizationFailedErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	videoList, err := service.NewQueryVideoService(ctx).QueryVideo(req, tokenUser)
	if err != nil {
		errMsg := err.Error()
		resp = &publish.PublishListResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	resp = &publish.PublishListResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, VideoList: videoList}
	return resp, nil
}

func getUserFromToken(token string, ctx context.Context) (*user.User, error) {
	claims, err := middleware.ParseToken(token)
	if err != nil {
		return nil, err
	}
	resp, err := rpc.User(ctx, &user.UserRequest{UserId: claims.ID, Token: token})
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}
