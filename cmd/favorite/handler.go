package main

import (
	"context"
	"tiktok-server/cmd/favorite/service"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
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

	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		errStr := err.Error()
		resp = &favorite.FavoriteActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	actionService := service.NewFavoriteActionService(ctx)
	if req.ActionType == 1 {
		err = actionService.AddFavorite(req, claims.ID)
	} else if req.ActionType == 2 {
		err = actionService.CancelFavorite(req, claims.ID)
	} else {
		err = erren.TypeNotSupportErr
	}

	if err != nil {
		errStr := err.Error()
		resp = &favorite.FavoriteActionResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &favorite.FavoriteActionResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg}
	return resp, nil
}

// ListFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) ListFavorite(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.UserId <= 0 {
		resp = &favorite.FavoriteListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	favorites, err := service.NewFavoriteListService(ctx).GetFavoriteList(req)
	if err != nil {
		errStr := err.Error()
		resp = &favorite.FavoriteListResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &favorite.FavoriteListResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, VideoList: favorites}
	return resp, nil
}

// MCheckFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MCheckFavorite(ctx context.Context, req *favorite.MCheckFavoriteRequest) (resp *favorite.MCheckFavoriteResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.UserId <= 0 || len(req.VideoIdList) == 0 {
		resp = &favorite.MCheckFavoriteResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	checkList, err := service.NewMCheckFavoriteService(ctx).MCheck(req)
	if err != nil {
		errStr := err.Error()
		resp = &favorite.MCheckFavoriteResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &favorite.MCheckFavoriteResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, ExistList: checkList}
	return resp, nil
}

// MCountFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) MCountFavorite(ctx context.Context, req *favorite.MCountVideoFavoriteRequest) (resp *favorite.MCountVideoFavoriteResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if len(req.VideoIdList) == 0 {
		resp = &favorite.MCountVideoFavoriteResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	countList, err := service.NewMCountFavoriteService(ctx).MCount(req)
	if err != nil {
		errStr := err.Error()
		resp = &favorite.MCountVideoFavoriteResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, err
	}

	resp = &favorite.MCountVideoFavoriteResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, CountList: countList}
	return resp, nil
}
