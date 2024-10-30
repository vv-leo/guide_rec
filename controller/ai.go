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
	prompt :=
		"# 要求:" + "\n" +
			"## 根据" + cueWord + "生成一个游玩的攻略" + "\n" +
			"## 生成的攻略规划必须包含标题，天数，每天游玩的景点及周边美食信息" + "\n" +
			"## 行程攻略尽可能简明扼要，不要出现任何的重复信息" + "\n" +
			"## 每天行程结束要推荐一个酒店,且酒店不能距离后一天第一个景点太远" + "\n" +
			"# 约束:" + "\n" +
			"## 文章标题不要超过30个字" + "\n" +
			"## 每天游玩的景点及周边美食信息不要超过5个"

	go utils.GetByDoubaoSSE("你是豆包，是一名出游规划专家", prompt, ch)
	for response := range ch {
		fmt.Fprintf(c.Writer, response)
		fmt.Printf(response)
		if f, ok := c.Writer.(http.Flusher); ok {
			f.Flush()
		}
		if len(response) == 0 || response == "end" {
			close(ch)
		}
	}
}
