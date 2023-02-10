package main

import (
	"context"
	"tiktok-server/cmd/feed/rpc"
	"tiktok-server/cmd/feed/service"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/feed"
	"tiktok-server/kitex_gen/user"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	resp = nil

	userId := int64(0)
	if req.Token != nil {
		tokenUser, err := getUserFromToken(*req.Token, ctx)
		if err != nil {
			errMsg := err.Error()
			resp = &feed.FeedResponse{StatusCode: erren.AuthorizationFailedErr.ErrCode, StatusMsg: &errMsg}
			return resp, err
		}
		userId = tokenUser.Id
	}

	videoList, nextTime, err := service.NewGetFeedService(ctx).GetFeedInfo(req, userId)
	if err != nil {
		errMsg := err.Error()
		resp = &feed.FeedResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errMsg}
		return resp, err
	}

	resp = &feed.FeedResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, VideoList: videoList, NextTime: &nextTime}
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
