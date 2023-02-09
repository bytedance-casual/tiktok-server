package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok-server/internal/conf"
	"tiktok-server/internal/erren"
	"tiktok-server/internal/middleware"
	"tiktok-server/kitex_gen/user"
	"tiktok-server/kitex_gen/user/userservice"
	"time"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		conf.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func User(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	resp, err := userClient.User(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func RegisterUser(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	resp, err := userClient.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func LoginUser(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	resp, err := userClient.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func MGetUsers(ctx context.Context, req *user.UsersMGetRequest) (*user.UsersMGetResponse, error) {
	resp, err := userClient.MGetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
