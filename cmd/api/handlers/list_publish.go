package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/publish"
)

type PublishListRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// ListPublish 查询登录用户的视频发布列表
func ListPublish(c *gin.Context) {
	var request PublishListRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.UserId <= 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ListPublish(c.Request.Context(), &publish.PublishListRequest{
		UserId: request.UserId,
		Token:  request.Token,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
