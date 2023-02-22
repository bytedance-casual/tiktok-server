package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok-server/cmd/relation/dal/db"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/relation"
)

type ActionRelationService struct {
	ctx context.Context
}

func NewActionRelationService(ctx context.Context) *ActionRelationService {
	return &ActionRelationService{ctx: ctx}
}

func (s *ActionRelationService) Follow(req *relation.RelationActionRequest) (err error) {
	claims, err := middleware.ParseToken(req.Token)
	if err != nil {
		return err
	}
	actor := strconv.FormatInt(claims.ID, 10)
	toUser := strconv.FormatInt(req.ToUserId, 10)
	if actor == toUser {
		return errors.New("你不能关注自己")
	}

	follow, err := db.QueryHasFollow(s.ctx, actor, toUser)
	if err != nil {
		return err
	}

	switch req.ActionType {
	case 1: //关注
		if follow == nil {
			return db.CreateFollow(s.ctx, actor, toUser)
		}
		if follow.IsFollow {
			return errors.New("你已经关注该用户")
		}
		err = db.UpdateFollow(s.ctx, follow.ID, true)
	case 2: //取消关注
		if follow == nil || follow.IsFollow == db.Canceled {
			return errors.New("未关注该用户")
		}
		err = db.UpdateFollow(s.ctx, follow.ID, false)
	default:
		err = errors.New("不支持的操作")
	}
	return
}
