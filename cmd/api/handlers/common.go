package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok-server/internal/erren"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// BadResponse 返回错误状态码与错误信息
func BadResponse(c *gin.Context, err error) {
	Err := erren.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

// SendResponse 返回数据
func SendResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
