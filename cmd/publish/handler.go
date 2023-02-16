package main

import (
	"context"
	"tiktok-server/cmd/publish/rpc"
	"tiktok-server/cmd/publish/service"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// ActionPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ActionPublish(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if len(req.Data) == 0 || len(req.Title) == 0 {
		resp = &publish.PublishActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, err
	}

	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		errMsg := err.Error()
		resp = &publish.PublishActionResponse{StatusCode: erren.AuthorizationFailedErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	err = service.NewUploadVideoService(ctx).UploadVideo(req, claims.ID)
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

	if req.UserId <= 0 {
		resp = &publish.PublishListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, err
	}

	tokenUser, err := rpc.GetUserFromToken(ctx, req.Token)
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

func (s *PublishServiceImpl) VideoActionPublish(ctx context.Context, req *publish.PublishVideoActionRequest) (resp *publish.PublishVideoActionResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.VideoId <= 0 {
		resp = &publish.PublishVideoActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, err
	}

	actionVideoService := service.NewActionVideoService(ctx)
	if req.ActionType == 1 {
		err = actionVideoService.UpdateFavorite(req)
	} else if req.ActionType == 2 {
		err = actionVideoService.UpdateComment(req)
	} else {
		err = erren.TypeNotSupportErr
	}

	if err != nil {
		errStr := err.Error()
		resp = &publish.PublishVideoActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &publish.PublishVideoActionResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg}
	return resp, nil
}