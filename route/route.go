package route

import (
	"github.com/gin-gonic/gin"
	"web3/controller"
)

func Init() {
	route := gin.Default()

	route.GET("/ai/generateGuide", func(c *gin.Context) {
		controller.GenerateGuide(c)
	})

	route.GET("/guide/create", func(c *gin.Context) {
		controller.Create(c)
	})

	route.Run(":8080")
}
