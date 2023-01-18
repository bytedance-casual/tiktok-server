package router

import (
	"github.com/gin-gonic/gin"
	"tiktok-server/internal/cont"
)

func Init(engine *gin.Engine) {
	engine.Static("/static", "../../web")
	router := engine.Group("/douyin")

	// basic apis
	router.GET("/feed/", cont.Feed)
	router.POST("/user/register/", cont.UserRegister)
	router.POST("/user/login/", cont.UserLogin)
	router.GET("/user/", cont.User)
	router.POST("/publish/action/", cont.PublishAction)
	router.GET("/publish/list/", cont.PublishList)

	//	// extra apis - I
	//	router.POST("/favorite/action/", cont.FavoriteAction)
	//	router.GET("/favorite/list/", cont.FavoriteList)
	//	router.POST("/comment/action/", cont.CommentAction)
	//	router.GET("/comment/list/", cont.CommentList)
	//
	//	// extra apis - II
	//	router.POST("/relation/action/", cont.RelationAction)
	//	router.GET("/relation/follow/list/", cont.FollowList)
	//	router.GET("/relation/follower/list/", cont.FollowerList)
	//	router.GET("/relation/friend/list/", cont.FriendList)
	//	router.GET("/message/chat/", cont.MessageChat)
	//	router.POST("/message/action/", cont.MessageAction)
}
