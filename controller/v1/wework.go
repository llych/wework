package v1

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"wework/extend/wework"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type weworkReqMsg struct {
	Tos     string `json:"tos"`
	Content string `json:"content"`
	Toparty string `json:"toparty"`
	Msg     string `json:"msg"`
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}

func Wework(c *gin.Context) {
	var (
		tos       string
		toparty   string
		content   string
		msg       string
		err       error
		bodyBytes []byte
	)
	bodyBytes, _ = ioutil.ReadAll(c.Request.Body)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	msgBody := weworkReqMsg{}
	if err := c.ShouldBindJSON(&msgBody); err == nil {
		tos, content, msg = msgBody.Tos, msgBody.Content, msgBody.Msg
	} else {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		tos, content, msg = c.PostForm("tos"), c.PostForm("content"), c.PostForm("msg")

	}
	if content == "" {
		content = msg
	}
	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "内容不能为空"})
		return
	}
	if tos == "" && toparty == "" {
		toparty = "1"
	}
	if err = wework.TextMsg(content, tos, toparty); err != nil {
		log.Error().Msgf("微信发送失败: %s|tos: %s, toparty: %s, content: %s", err.Error(), tos, toparty, strings.ReplaceAll(content, "\n", "\\n"))
		c.JSON(http.StatusCreated, gin.H{"msg": "发送失败", "weworkErr": err.Error()})
		return
	}
	log.Info().Msgf("微信发送成功|tos: %s, toparty: %s, content: %s", tos, toparty, strings.ReplaceAll(content, "\n", "\\n"))
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})

}
