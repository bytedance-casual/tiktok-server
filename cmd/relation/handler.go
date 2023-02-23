package main

import (
	"context"
	"tiktok-server/cmd/relation/service"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// ActionRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ActionRelation(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = &relation.RelationActionResponse{
		StatusCode: erren.SuccessCode,
		StatusMsg:  nil,
	}
	if req.ToUserId <= 0 || len(req.Token) == 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp = &relation.RelationActionResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}
	err = service.NewActionRelationService(ctx).Follow(req)
	if err != nil {
		msg := err.Error()
		resp = &relation.RelationActionResponse{
			StatusCode: erren.ServiceErr.ErrCode,
			StatusMsg:  &msg,
		}
		return
	}
	return
}

// ListFollowRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollowRelation(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	resp = nil
	if req.UserId <= 0 || len(req.Token) == 0 {
		resp = &relation.RelationFollowListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}
	//fmt.Println("hander1")
	users, err := service.NewFollowListService(ctx).GetFollowList(req)
	//fmt.Println("hander2")
	if err != nil {
		errStr := err.Error()
		resp = &relation.RelationFollowListResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, nil
	}
	resp = &relation.RelationFollowListResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, UserList: users}
	return resp, nil
}

// ListFollowerRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollowerRelation(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = nil
	if req.UserId <= 0 || len(req.Token) == 0 {
		resp = &relation.RelationFollowerListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}
	users, err := service.NewFollowerListService(ctx).GetFollowerList(req)
	if err != nil {
		errStr := err.Error()
		resp = &relation.RelationFollowerListResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, nil
	}
	resp = &relation.RelationFollowerListResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, UserList: users}
	return resp, nil
}

// ListFriendRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFriendRelation(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	if req.UserId <= 0 || len(req.Token) == 0 {
		resp = &relation.RelationFriendListResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}
	claims, err := middleware.ParseToken(req.Token)
	friends, err := service.NewFriendListService(ctx).ListFriend(claims.ID)
	if err != nil {
		errStr := err.Error()
		resp = &relation.RelationFriendListResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, nil
	}
	resp = &relation.RelationFriendListResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, UserList: friends}
	return resp, nil

}

// MCheckFollowRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MCheckFollowRelation(ctx context.Context, req *relation.MCheckFollowRelationRequest) (resp *relation.MCheckFollowRelationResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if req.UserId <= 0 || len(req.UserIdList) == 0 {
		resp = &relation.MCheckFollowRelationResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	checkList, err := service.NewMCheckFollowRelationService(ctx).MCheckFollow(req)
	if err != nil {
		errStr := err.Error()
		resp = &relation.MCheckFollowRelationResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, nil
	}

	resp = &relation.MCheckFollowRelationResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, CheckList: checkList}
	return resp, nil
}

// MCountRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MCountRelation(ctx context.Context, req *relation.MCountRelationRequest) (resp *relation.MCountRelationResponse, err error) {
	// TODO: Your code here...
	resp = nil

	if len(req.UserIdList) == 0 {
		resp = &relation.MCountRelationResponse{StatusCode: erren.ParamErr.ErrCode, StatusMsg: &erren.ParamErr.ErrMsg}
		return resp, nil
	}

	followCountList, followerCountList, err := service.NewMCountRelationService(ctx).MCount(req)
	if err != nil {
		errStr := err.Error()
		resp = &relation.MCountRelationResponse{StatusCode: erren.ServiceErr.ErrCode, StatusMsg: &errStr}
		return resp, nil
	}

	resp = &relation.MCountRelationResponse{StatusCode: erren.SuccessCode, StatusMsg: &erren.Success.ErrMsg, FollowCountList: followCountList, FollowerCountList: followerCountList}
	return resp, nil
}
