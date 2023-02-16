package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/erren"
	"tiktok-server/kitex_gen/user"
)

type UserRegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var request UserRegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		BadResponse(c, err)
		return
	}

	fmt.Println(request)
	if len(request.Username) == 0 || len(request.Password) == 0 {
		BadResponse(c, erren.ParamErr)
		return
	}

	resp, err := rpc.RegisterUser(c.Request.Context(), &user.UserRegisterRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		BadResponse(c, erren.ConvertErr(err))
		return
	}
	SendResponse(c, resp)
}
