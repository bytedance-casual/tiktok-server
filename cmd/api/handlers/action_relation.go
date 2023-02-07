package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/relation"
)

var RelationActionType = map[int32]struct{}{1: {}, 2: {}}

// ActionRelation 登录用户对其他用户进行关注或取消关注
func ActionRelation(c *gin.Context) {
	var request relation.RelationActionRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	_, exist := RelationActionType[request.ActionType]
	if request.ToUserId <= 0 || !exist {
		BadResponse(c, erren.ParamErr)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.ActionRelation(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
