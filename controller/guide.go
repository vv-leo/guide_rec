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

func GuideList(c *gin.Context) {
	owner := c.Query("owner")
	if len(owner) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "owner not null"})
		return
	}
	res, err := guideSer.ListByOwner(owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GuideDetail(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id not null"})
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
	from := c.Query("from")
	status := c.Query("status")
	if len(id) == 0 || len(from) == 0 || len(status) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "param error"})
		return
	}
	success, err := guideSer.ThumbsUp(id, from, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": success})
}

func GuideRec(c *gin.Context) {
	res, err := guideSer.Rec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
