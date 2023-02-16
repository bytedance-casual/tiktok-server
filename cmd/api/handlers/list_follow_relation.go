package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/relation"
)

type RelationFollowListRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// ListFollowRelation 登录用户关注的所有用户列表
func ListFollowRelation(c *gin.Context) {
	var request RelationFollowListRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.UserId <= 0 || len(request.Token) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ListFollowRelation(c.Request.Context(), &relation.RelationFollowListRequest{
		UserId: request.UserId,
		Token:  request.Token,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
