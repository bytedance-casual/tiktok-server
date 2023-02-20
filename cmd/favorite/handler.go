package main

import (
	"context"
	"tiktok-server/cmd/favorite/rpc"
	"tiktok-server/cmd/favorite/service"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// ActionFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) ActionFavorite(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = nil
	if req.VideoId <= 0 {
		resp = &favorite.FavoriteActionResponse{StatusCode: erren.ParamErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return nil, err
	}
	if req.ActionType == 1 {
		resp, err = doAddFavorite(ctx, req)
	} else if req.ActionType == 2 {
		resp, err = doCancelFavorite(ctx, req)
	} else {
		err = erren.TypeNotSupportErr
	}
	return
}
func doAddFavorite(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = nil
	if req.UserId <= 0 || req.VideoId <= 0 {
		resp = &favorite.FavoriteActionResponse{StatusCode: erren.ParamErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return nil, err
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		errStr := err.Error()
		resp = &favorite.FavoriteActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}
	resp = &favorite.FavoriteActionResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg}
	return resp, err
}
func doCancelFavorite(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = nil
	if req.UserId <= 0 || req.VideoId <= 0 {
		resp = &favorite.FavoriteActionResponse{StatusCode: erren.ParamErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return nil, err
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		errStr := err.Error()
		resp = &favorite.FavoriteActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}
	resp = &favorite.FavoriteActionResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg}
	return resp, err
}

// ListFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) ListFavorite(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = nil
	if req.UserId <= 0 {
		resp = &favorite.FavoriteListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}
	user, err := rpc.GetUserFromToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	favorites, err := service.NewFavoriteListService(ctx).GetFavoriteList(req, user)
	if err != nil {
		errStr := err.Error()
		resp = &favorite.FavoriteListResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &favorite.FavoriteListResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, VideoList: favorites}
	return resp, nil
}
