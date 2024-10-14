package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

/*
*
SSE生成行程规划
生成最多20s，超过20s也不生成了
*/
func GenerateGuideSSE(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	start := time.Now()

	cueWord := c.Query("cueWord")
	ch := make(chan string)
	go utils.GetByDoubaoSSE("你是豆包，是一名出游规划专家", cueWord+"，必须控制在500字以内", ch)
	for response := range ch {
		fmt.Fprintf(c.Writer, response)
		if f, ok := c.Writer.(http.Flusher); ok {
			f.Flush()
			//超过20关闭生成通道
			if time.Since(start).Milliseconds() > 20000 {
				break
			}
		}
	}
}
