package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web3/service"
)

var buyRecordService = service.NewBuyRecordService()

func BuyRecordListByGuideId(c *gin.Context) {
	guideId := c.Query("guide_id")
	if len(guideId) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "guide_id not null"})
		return
	}
	res, err := buyRecordService.GetByGuideId(guideId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})

}
