package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/user"
)

type UserLoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// LoginUser 用户登录
func LoginUser(c *gin.Context) {
	var request UserLoginRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if len(request.Username) == 0 || len(request.Password) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.LoginUser(c.Request.Context(), &user.UserLoginRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
