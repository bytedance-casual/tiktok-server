package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tiktok-server/internal/conf"
	"tiktok-server/internal/db"
	"tiktok-server/internal/router"
)

func main() {
	err := conf.Init()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}
	engine := gin.Default()
	router.Init(engine)
	err = engine.Run(conf.Config.Server.Address)
	if err != nil {
		log.Fatal(err)
	}
}
