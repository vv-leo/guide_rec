package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web3/model"
	"web3/service"
)

var userSer = service.NewUserService()

func UserCreate(c *gin.Context) {
	var user *model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success, err := userSer.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": success})

}

func UserDetail(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id不能为空"})
		return
	}
	user, err := userSer.Detail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UserFollow(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	status := c.Query("status")
	if len(from) == 0 || len(to) == 0 || len(status) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "参数错误"})
		return
	}
	success, err := userSer.Follow(from, to, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": success})
}
