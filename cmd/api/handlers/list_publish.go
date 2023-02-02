package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/publish"
)

// ListPublish 查询登录用户的视频发布列表
func ListPublish(c *gin.Context) {
	var request publish.PublishListRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.UserId <= 0 || len(request.Token) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.ListPublish(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
