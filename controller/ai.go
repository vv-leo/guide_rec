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

	prompt := "#任务:" + "\n" +
		"'''" + "\n" +
		cueWord + "\n" +
		"'''" + "\n" +
		"# 要求:" + "\n" +
		"## 根据被''' '''包围的任务信息规划一条行程规划" + "\n" +
		"## 每天的行程需要有景点，美食，住宿信息" + "\n" +
		"## 简略写明推荐理由即可" + "\n" +
		"## 行程最后要写明出行建议等信息" + "\n" +
		"# 约束:" + "\n" +
		"## 每天推荐至少2个景点且不能重复" + "\n" +
		"## 每天推荐的美食需要在景点附近" + "\n" +
		"## 输出的内容中不要出现特殊字符，如：#$%&等" + "\n" +
		"# 例子: 自由发挥，尽量美观即可"

	go utils.GetByDoubaoSSE("你是豆包，是一名出游规划专家", prompt, ch)
	for response := range ch {
		fmt.Fprintf(c.Writer, response)
		//fmt.Printf(response)
		if f, ok := c.Writer.(http.Flusher); ok {
			f.Flush()
		}
	}
}
