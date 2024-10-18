package route

import (
	"github.com/gin-gonic/gin"
	"web3/controller"
)

func Init() {
	route := gin.Default()

	route.GET("/ai/generateGuide", func(c *gin.Context) { controller.GenerateGuide(c) })

	route.GET("/ai/generateGuideSSE", func(c *gin.Context) { controller.GenerateGuideSSE(c) })

	route.POST("/guide/create", func(c *gin.Context) { controller.GuideCreate(c) })

	route.GET("/guide/detail", func(c *gin.Context) { controller.GuideDetail(c) })

	route.GET("/guide/thumbsUp", func(c *gin.Context) { controller.GuideThumbsUp(c) })

	route.Run(":8080")
}
