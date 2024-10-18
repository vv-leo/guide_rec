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
		return
	}

	success, err := guideSer.Create(guide)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": success})
}

func GuideDetail(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id不能为空"})
		return
	}
	guide, err := guideSer.Detail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": guide})
}

func GuideThumbsUp(c *gin.Context) {
	id := c.Query("id")
	account := c.Query("account")
	status := c.Query("status")
	if len(id) == 0 || len(account) == 0 || len(status) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "参数错误"})
		return
	}
	guide, err := guideSer.ThumbsUp(id, account, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": guide})
}
