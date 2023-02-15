package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/feed"
)

type FeedRequest struct {
	LatestTime *int64  `form:"latest_time"`
	Token      *string `form:"token"`
}

// Feed 不限制登录，返回视频流，视频数由服务端控制，单次最多30个。
func Feed(c *gin.Context) {
	var request FeedRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	resp, err := rpc.Feed(c.Request.Context(), &feed.FeedRequest{
		LatestTime: request.LatestTime,
		Token:      request.Token,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
