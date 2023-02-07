package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/user"
)

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var request user.UserRegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}
	//fmt.Println(c.Query("username"))
	fmt.Println(request)
	if len(request.Username) == 0 || len(request.Password) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	// gin 貌似没有配套上下文参数，暂时手动创建
	ctx := context.Background()
	resp, err := rpc.RegisterUser(ctx, &request)
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
