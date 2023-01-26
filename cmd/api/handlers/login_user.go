package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/user"
)

// LoginUser 用户登录
func LoginUser(c *gin.Context) {
	var request user.UserLoginRequest
	if err := c.Bind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	if len(request.Username) == 0 || len(request.Password) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.UserLogin(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
