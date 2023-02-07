package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/favorite"
)

var FavoriteActionType = map[int32]struct{}{1: {}, 2: {}}

// ActionFavorite 登录用户对视频的点赞和取消点赞操作
func ActionFavorite(c *gin.Context) {
	var request favorite.FavoriteActionRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	_, exist := FavoriteActionType[request.ActionType]
	if request.VideoId <= 0 || !exist {
		BadResponse(c, erren.ParamErr)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.ActionFavorite(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
