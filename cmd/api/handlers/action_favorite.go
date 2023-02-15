package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/favorite"
)

type FavoriteActionRequest struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int32  `form:"action_type"`
}

var FavoriteActionType = map[int32]struct{}{1: {}, 2: {}}

// ActionFavorite 登录用户对视频的点赞和取消点赞操作
func ActionFavorite(c *gin.Context) {
	var request FavoriteActionRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	_, exist := FavoriteActionType[request.ActionType]
	if request.VideoId <= 0 || !exist {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ActionFavorite(c.Request.Context(), &favorite.FavoriteActionRequest{
		Token:      request.Token,
		VideoId:    request.VideoId,
		ActionType: request.ActionType,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
