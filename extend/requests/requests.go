package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

var client http.Client

type req struct {
	client http.Client
}

func (r *req) Post(url string, body []byte) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if resp, err = r.client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody, nil
}

func (r *req) Get(url string) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	if resp, err = r.client.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody, nil
}

func InitReq() *req {
	req := &req{}
	jar, _ := cookiejar.New(nil)
	req.client.Jar = jar
	return req
}
