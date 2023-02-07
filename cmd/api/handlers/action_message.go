package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/message"
)

var MessageActionType = map[int32]struct{}{1: {}}

// ActionMessage 登录用户对消息的相关操作，目前只支持消息发送
func ActionMessage(c *gin.Context) {
	var request message.MessageActionRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	_, exist := MessageActionType[request.ActionType]
	if request.ToUserId <= 0 || len(request.Content) == 0 || !exist {
		BadResponse(c, erren.ParamErr)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.ActionMessage(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
