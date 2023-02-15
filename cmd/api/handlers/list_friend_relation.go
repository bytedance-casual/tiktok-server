package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/relation"
)

type RelationFriendListRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// ListFriendRelation 所有关注登录用户的粉丝列表
func ListFriendRelation(c *gin.Context) {
	var request RelationFriendListRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.UserId <= 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ListFriendRelation(c.Request.Context(), &relation.RelationFriendListRequest{
		UserId: request.UserId,
		Token:  request.Token,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
