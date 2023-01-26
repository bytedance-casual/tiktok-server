package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tiktok-server/cmd/api/rpc"
	"tiktok-server/internal/conf"
	"tiktok-server/internal/tracer"
)

func Init() {
	tracer.InitJaeger(conf.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	engine := gin.Default()
	initRouter(engine)
	err := engine.Run(conf.Config.Server.Address)
	if err != nil {
		log.Fatal(err)
	}
}
