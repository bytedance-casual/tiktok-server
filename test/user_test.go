package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/kitex_gen/user"
)

var uuidName string

func TestUser(t *testing.T) {
	resp, err := rpc.User(ctx, &user.UserRequest{
		UserId: 3,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestRegisterUser(t *testing.T) {
	uuidName = uuid.NewV4().String()
	resp, err := rpc.RegisterUser(ctx, &user.UserRegisterRequest{
		Username: uuidName,
		Password: "123456",
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestLoginUser(t *testing.T) {
	resp, err := rpc.LoginUser(ctx, &user.UserLoginRequest{
		Username: uuidName,
		Password: "123456",
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func TestMGetUsers(t *testing.T) {
	resp, err := rpc.MGetUsers(ctx, &user.UsersMGetRequest{
		UserId:     2,
		UserIdList: []int64{3},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
