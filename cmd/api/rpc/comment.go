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
	"tiktok-server/kitex_gen/comment"
	"tiktok-server/kitex_gen/comment/commentservice"
	"time"
)

var commentClient commentservice.Client

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := commentservice.NewClient(
		conf.CommentServiceName,
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
	commentClient = c
}

func ActionComment(ctx context.Context, req *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	resp, err := commentClient.ActionComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func ListComment(ctx context.Context, req *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	resp, err := commentClient.ListComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
