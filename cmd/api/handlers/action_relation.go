package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/relation"
)

type RelationActionRequest struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
}

var RelationActionType = map[int32]struct{}{1: {}, 2: {}}

// ActionRelation 登录用户对其他用户进行关注或取消关注
func ActionRelation(c *gin.Context) {
	var request RelationActionRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	_, exist := RelationActionType[request.ActionType]
	if request.ToUserId <= 0 || !exist {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ActionRelation(c.Request.Context(), &relation.RelationActionRequest{
		Token:      request.Token,
		ToUserId:   request.ToUserId,
		ActionType: request.ActionType,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
