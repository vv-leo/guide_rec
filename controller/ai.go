package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web3/service"
)

var aiSer = service.NewAiService()

func GenerateGuide(c *gin.Context) {
	cueWord := c.Query("cueWord")
	res, err := aiSer.GenerateGuide(cueWord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "网络超时"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
