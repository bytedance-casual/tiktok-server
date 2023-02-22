package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
	"tiktok-server/cmd/user/dal"
	"tiktok-server/cmd/user/rpc"
	"tiktok-server/internal/bound"
	"tiktok-server/internal/conf"
	"tiktok-server/internal/middleware"
	"tiktok-server/internal/tracer"
	user "tiktok-server/kitex_gen/user/userservice"
)

func Init() {
	conf.Init()
	tracer.InitJaeger(conf.UserServiceName)
	dal.Init()
	rpc.InitRPC()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8886")
	if err != nil {
		panic(err)
	}
	Init()
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.UserServiceName}), // server name
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
