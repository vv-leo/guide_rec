package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web3/model"
	"web3/service"
)

var guideSer = service.NewGuideService()

func GuideCreate(c *gin.Context) {
	var guide model.Guide
	err := c.ShouldBindJSON(&guide)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	success, err := guideSer.Create(guide)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": success})
}

func GuideDetail(c *gin.Context) {
	id := c.GetString("id")
	guideSer.Detail(id)
}
