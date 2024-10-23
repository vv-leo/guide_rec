package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"web3/controller"
	"web3/ether/call"
)

func Init() {

	route := gin.Default()

	route.Use(cors.Default())

	route.GET("/ai/generateGuide", func(c *gin.Context) { controller.GenerateGuide(c) })

	route.GET("/ai/generateGuideSSE", func(c *gin.Context) { controller.GenerateGuideSSE(c) })

	route.POST("/guide/create", func(c *gin.Context) { controller.GuideCreate(c) })

	route.GET("/guide/detail", func(c *gin.Context) { controller.GuideDetail(c) })

	route.GET("/guide/rec", func(c *gin.Context) { controller.GuideRec(c) })

	route.GET("/guide/listByOwner", func(c *gin.Context) { controller.GuideList(c) })

	route.GET("/guide/thumbsUp", func(c *gin.Context) { controller.GuideThumbsUp(c) })

	route.POST("/user/create", func(c *gin.Context) { controller.UserCreate(c) })

	route.GET("/user/detail", func(c *gin.Context) { controller.UserDetail(c) })

	route.GET("/user/follow", func(c *gin.Context) { controller.UserFollow(c) })

	route.GET("/user/thumbsUped", func(c *gin.Context) { controller.UserThumbsUped(c) })

	route.GET("/user/followed", func(c *gin.Context) { controller.UserFollowed(c) })

	route.GET("/test/ListNFT", func(c *gin.Context) { call.ListNFT() })

	route.Run(":8080")
}
