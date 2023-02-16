package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/message"
)

type MessageActionRequest struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
	Content    string `form:"content"`
}

var MessageActionType = map[int32]struct{}{1: {}}

// ActionMessage 登录用户对消息的相关操作，目前只支持消息发送
func ActionMessage(c *gin.Context) {
	var request MessageActionRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	_, exist := MessageActionType[request.ActionType]
	if request.ToUserId <= 0 || len(request.Content) == 0 || !exist {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ActionMessage(c.Request.Context(), &message.MessageActionRequest{
		Token:      request.Token,
		ToUserId:   request.ToUserId,
		ActionType: request.ActionType,
		Content:    request.Content,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
