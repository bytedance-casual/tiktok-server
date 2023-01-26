package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/publish"
)

// ActionPublish 登录用户视频投稿
func ActionPublish(c *gin.Context) {
	var request publish.PublishActionRequest
	if err := c.Bind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if len(request.Token) == 0 || len(request.Title) == 0 || len(request.Data) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.ActionPublish(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
