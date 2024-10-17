package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web3/service"
	"web3/utils"
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

func GenerateGuideSSE(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	cueWord := c.Query("cueWord")
	ch := make(chan string)

	go utils.GetByDoubaoSSE("你是豆包，是一名出游规划专家", cueWord+"，必须控制在150字以内", ch)
	for response := range ch {
		fmt.Fprintf(c.Writer, response)
		fmt.Printf(response)
		if f, ok := c.Writer.(http.Flusher); ok {
			f.Flush()
		}
	}
}
