package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/message"
)

type MessageChatRequest struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	PreMsgTime int64  `form:"pre_msg_time"`
}

// ChatMessage 当前登录用户和其他指定用户的聊天消息记录
func ChatMessage(c *gin.Context) {
	var request MessageChatRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.ToUserId <= 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ChatMessage(c.Request.Context(), &message.MessageChatRequest{
		Token:      request.Token,
		ToUserId:   request.ToUserId,
		PreMsgTime: request.PreMsgTime,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
