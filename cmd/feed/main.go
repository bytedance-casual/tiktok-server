package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
	"tiktok-server/cmd/feed/dal"
	"tiktok-server/cmd/feed/rpc"
	"tiktok-server/internal/bound"
	"tiktok-server/internal/conf"
	"tiktok-server/internal/middleware"
	"tiktok-server/internal/tracer"
	feed "tiktok-server/kitex_gen/feed/feedservice"
)

func Init() {
	conf.Init()
	tracer.InitJaeger(conf.FeedServiceName)
	rpc.InitRPC()
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8882")
	if err != nil {
		panic(err)
	}
	Init()
	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.FeedServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                        // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
