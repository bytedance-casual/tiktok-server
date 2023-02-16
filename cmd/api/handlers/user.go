package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/user"
)

type UserRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// User 查询用户信息
func User(c *gin.Context) {
	var request UserRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if request.UserId <= 0 || len(request.Token) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.User(c.Request.Context(), &user.UserRequest{
		UserId: request.UserId,
		Token:  request.Token,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
