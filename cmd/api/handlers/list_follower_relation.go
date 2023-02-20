package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/relation"
)

type RelationFollowerListRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// ListFollowerRelation 所有关注登录用户的粉丝列表
func ListFollowerRelation(c *gin.Context) {
	var request RelationFollowerListRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.UserId <= 0 || len(request.Token) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.ListFollowerRelation(c.Request.Context(), &relation.RelationFollowerListRequest{
		UserId: request.UserId,
		Token:  request.Token,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
