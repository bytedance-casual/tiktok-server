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
	doUser(t)
}

func TestRegisterUser(t *testing.T) {
	doRegisterUser(t)
}

func TestLoginUser(t *testing.T) {
	doLoginUser(t)
}

func TestMGetUsers(t *testing.T) {
	doMGetUsers(t)
}

func BenchmarkUser(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doUser(b)
		}
	})
}

func BenchmarkRegisterUser(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doRegisterUser(b)
		}
	})
}

func BenchmarkLoginUser(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doLoginUser(b)
		}
	})
}

func BenchmarkMGetUsers(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doMGetUsers(b)
		}
	})
}

func doUser(t assert.TestingT) {
	resp, err := rpc.User(ctx, &user.UserRequest{
		UserId: 3,
		Token:  TOKEN,
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doRegisterUser(t assert.TestingT) {
	uuidName = uuid.NewV4().String()
	resp, err := rpc.RegisterUser(ctx, &user.UserRegisterRequest{
		Username: uuidName,
		Password: "123456",
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doLoginUser(t assert.TestingT) {
	resp, err := rpc.LoginUser(ctx, &user.UserLoginRequest{
		Username: uuidName,
		Password: "123456",
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}

func doMGetUsers(t assert.TestingT) {
	resp, err := rpc.MGetUsers(ctx, &user.UsersMGetRequest{
		UserId:     2,
		UserIdList: []int64{3},
	})
	assert.NoError(t, err)
	fmt.Printf("%v", resp)
}
