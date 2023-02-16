package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/comment"
)

type CommentActionRequest struct {
	Token       string  `form:"token"`
	VideoId     int64   `form:"video_id"`
	ActionType  int32   `form:"action_type"`
	CommentText *string `form:"comment_text"`
	CommentId   *int64  `form:"comment_id"`
}

var CommentActionType = map[int32]struct{}{1: {}, 2: {}}

// ActionComment 登录用户对视频进行评论
func ActionComment(c *gin.Context) {
	var request CommentActionRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	_, exist := CommentActionType[request.ActionType]
	if request.VideoId <= 0 || !exist {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ActionComment(c.Request.Context(), &comment.CommentActionRequest{
		Token:       request.Token,
		VideoId:     request.VideoId,
		ActionType:  request.ActionType,
		CommentText: request.CommentText,
		CommentId:   request.CommentId,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
