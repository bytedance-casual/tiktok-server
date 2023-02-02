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
	"tiktok-server/kitex_gen/feed"
	"tiktok-server/kitex_gen/feed/feedservice"
	"time"
)

var feedClient feedservice.Client

func initFeedRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := feedservice.NewClient(
		conf.FeedServiceName,
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
	feedClient = c
}

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	resp, err := feedClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
