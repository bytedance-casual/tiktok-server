package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/feed"
)

// Feed 不限制登录，返回视频流
func Feed(c *gin.Context) {
	var request feed.FeedRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.Feed(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
