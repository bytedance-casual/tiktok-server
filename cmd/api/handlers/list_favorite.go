package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/favorite"
)

type FavoriteListRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// ListFavorite 登录用户的所有点赞视频
func ListFavorite(c *gin.Context) {
	var request FavoriteListRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.UserId <= 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ListFavorite(c.Request.Context(), &favorite.FavoriteListRequest{
		UserId: request.UserId,
		Token:  request.Token,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
