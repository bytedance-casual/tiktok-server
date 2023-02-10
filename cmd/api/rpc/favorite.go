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
	"tiktok-server/kitex_gen/favorite"
	"tiktok-server/kitex_gen/favorite/favoriteservice"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := favoriteservice.NewClient(
		conf.FavoriteServiceName,
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
	favoriteClient = c
}

func ActionFavorite(ctx context.Context, req *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	resp, err := favoriteClient.ActionFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func ListFavorite(ctx context.Context, req *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	resp, err := favoriteClient.ListFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
