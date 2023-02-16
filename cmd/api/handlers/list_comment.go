package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/comment"
)

type CommentListRequest struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id"`
}

// ListComment 查看视频的所有评论，按发布时间倒序
func ListComment(c *gin.Context) {
	var request CommentListRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.VideoId <= 0 || len(request.Token) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ListComment(c.Request.Context(), &comment.CommentListRequest{
		Token:   request.Token,
		VideoId: request.VideoId,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
