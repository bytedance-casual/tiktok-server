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
	"tiktok-server/kitex_gen/relation"
	"tiktok-server/kitex_gen/relation/relationservice"
	"time"
)

var relationClient relationservice.Client

func initRelationRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := relationservice.NewClient(
		conf.RelationServiceName,
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
	relationClient = c
}

func ActionRelation(ctx context.Context, req *relation.RelationActionRequest) (*relation.RelationActionResponse, error) {
	resp, err := relationClient.ActionRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func ListFollowRelation(ctx context.Context, req *relation.RelationFollowListRequest) (*relation.RelationFollowListResponse, error) {
	resp, err := relationClient.ListFollowRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func ListFollowerRelation(ctx context.Context, req *relation.RelationFollowerListRequest) (*relation.RelationFollowerListResponse, error) {
	resp, err := relationClient.ListFollowerRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func ListFriendRelation(ctx context.Context, req *relation.RelationFriendListRequest) (*relation.RelationFriendListResponse, error) {
	resp, err := relationClient.ListFriendRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func MCheckFollowRelation(ctx context.Context, req *relation.MCheckFollowRelationRequest) (*relation.MCheckFollowRelationResponse, error) {
	resp, err := relationClient.MCheckFollowRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

func MCountRelation(ctx context.Context, req *relation.MCountRelationRequest) (*relation.MCountRelationResponse, error) {
	resp, err := relationClient.MCountRelation(ctx, req)
	if err != nil {
		return nil, err
	}
	if _, ok := erren.ErrorMap[resp.StatusCode]; ok {
		return nil, erren.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
