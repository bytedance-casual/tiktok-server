package main

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/cmd/api/handlers"
)

func initRouter(engine *gin.Engine) {
	engine.Static("/static", "../../web")
	router := engine.Group("/douyin")

	// basic apis
	router.GET("/feed/", handlers.Feed)
	router.POST("/user/register/", handlers.RegisterUser)
	router.POST("/user/login/", handlers.LoginUser)
	router.GET("/user/", handlers.User)
	router.POST("/publish/action/", handlers.ActionPublish)
	router.GET("/publish/list/", handlers.ListPublish)

	//	// extra apis - I
	//	router.POST("/favorite/action/", handlers.FavoriteAction)
	//	router.GET("/favorite/list/", handlers.FavoriteList)
	//	router.POST("/comment/action/", handlers.CommentAction)
	//	router.GET("/comment/list/", handlers.CommentList)
	//
	//	// extra apis - II
	//	router.POST("/relation/action/", handlers.RelationAction)
	//	router.GET("/relation/follow/list/", handlers.FollowList)
	//	router.GET("/relation/follower/list/", handlers.FollowerList)
	//	router.GET("/relation/friend/list/", handlers.FriendList)
	//	router.GET("/message/chat/", handlers.MessageChat)
	//	router.POST("/message/action/", handlers.MessageAction)
}
