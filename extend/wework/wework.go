package wework

import (
	"errors"
	"fmt"
	"time"
	"wework/extend/cache"
	"wework/extend/conf"
	"wework/extend/requests"

	"github.com/rs/zerolog/log"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type weworkToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type weworkResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	Invaliduser string `json:"invaliduser"`
}

func Token() (string, error) {
	//从缓存中获取token
	if token, ok := cache.Cache.Get("token"); ok {
		return token.(string), nil
	}
	var (
		url       string
		err       error
		respBody  []byte
		tokenResp *weworkToken
	)
	req := requests.InitReq()

	url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		conf.WeworkConf.CorpID,
		conf.WeworkConf.Secret,
	)

	if respBody, err = req.Get(url); err != nil {
		return "", errors.New("请求token失败")
	}
	tokenResp = &weworkToken{}
	if err = json.Unmarshal(respBody, tokenResp); err != nil {
		return "", errors.New("解析请求失败")
	}
	if tokenResp.Errcode == 0 {
		cache.Cache.Set("token", tokenResp.AccessToken, time.Hour*1)
	}
	return tokenResp.AccessToken, nil
}

func RetryToken() (string, error) {
	cache.Cache.Delete("token")
	return Token()
}

func ClearToken() {
	cache.Cache.Delete("token")
}

func TextMsg(msg string, tos string, toparty string) error {
	var (
		url      string
		err      error
		token    string
		postBody []byte
		resp     []byte
	)
	for i := 0; i < 3; i++ {
		if token, err = Token(); err != nil {
			return err
		}
		url = fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)

		postData := map[string]interface{}{
			"touser":  tos,
			"toparty": toparty,
			"msgtype": "text",
			"agentid": conf.WeworkConf.AgentId,
			"text": map[string]string{
				"content": msg,
			},
			"safe":            0,
			"enable_id_trans": 0,
		}
		if postBody, err = json.Marshal(postData); err != nil {
			log.Error().Msgf("发送请求失败, 尝试重试 %d", i+1)
			ClearToken()
			continue
		}

		req := requests.InitReq()
		if resp, err = req.Post(url, postBody); err != nil {
			return err
		}
		wResp := &weworkResponse{}
		if err = json.Unmarshal(resp, wResp); err != nil {
			return errors.New("解析请求包出错")
		}
		if wResp.Errcode != 0 {
			return errors.New(wResp.Errmsg)
		}
		if len(wResp.Invaliduser) > 0 {
			return fmt.Errorf("Invaliduser: %s", wResp.Invaliduser)
		}
		return nil
	}

	return err
}
