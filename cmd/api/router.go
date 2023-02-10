package main

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/handlers"
	"tiktok-server/internal/middleware"
)

func initRouter(engine *gin.Engine) {
	engine.Static("/static", "../../web")
	router := engine.Group("/douyin")

	// basic apis
	router.POST("/user/register/", handlers.RegisterUser)
	router.POST("/user/login/", handlers.LoginUser)
	router.GET("/feed/", handlers.Feed)

	// jwt middleware
	router.Use(middleware.JWTAuthMiddleware())

	// 后续接口经由 jwt 校验，无需再校验 Token 字段
	router.GET("/user/", handlers.User)
	router.POST("/publish/action/", handlers.ActionPublish)
	router.GET("/publish/list/", handlers.ListPublish)

	// extra apis - I
	router.POST("/favorite/action/", handlers.ActionFavorite)
	router.GET("/favorite/list/", handlers.ListFavorite)
	router.POST("/comment/action/", handlers.ActionComment)
	router.GET("/comment/list/", handlers.ListComment)

	// extra apis - II
	router.POST("/relation/action/", handlers.ActionRelation)
	router.GET("/relation/follow/list/", handlers.ListFollowRelation)
	router.GET("/relation/follower/list/", handlers.ListFollowerRelation)
	router.GET("/relation/friend/list/", handlers.ListFriendRelation)
	router.GET("/message/chat/", handlers.ChatMessage)
	router.POST("/message/action/", handlers.ActionMessage)
}
