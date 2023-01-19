package cont

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok-server/internal/model"
	"tiktok-server/internal/service"
)

func UserRegister(c *gin.Context) {
	response, err := service.RegisterUser(&model.UserRegisterRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	})
	if err != nil {
		c.String(http.StatusBadRequest, "register failed:%s", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func UserLogin(c *gin.Context) {

}

func User(c *gin.Context) {

}
