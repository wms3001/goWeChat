package goWeChat

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/wms3001/goWeChat/model"
	"log"
)

type GoWeChat struct {
	Url          string `json:"url"`
	Corpid       string `json:"corpid"`
	Corpsecret   string `json:"corpsecret"`
	Access_token string `json:"access_Token"`
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	log.SetPrefix("Message:")
}

func (goWeChat *GoWeChat) GetAccessToken() model.AccessToken {
	url := goWeChat.Url + "/cgi-bin/gettoken?corpid=" + goWeChat.Corpid + "&corpsecret=" + goWeChat.Corpsecret
	res := goWeChat.Get(url)
	accessToken := model.AccessToken{}
	json.Unmarshal([]byte(res), &accessToken)
	goWeChat.Access_token = accessToken.AccessToken
	return accessToken
}

func (goWeChat *GoWeChat) SendTextMessage(message model.TextMessage) model.TextMessageResponse {
	url := goWeChat.Url + "/cgi-bin/message/send?access_token=" + goWeChat.Access_token
	req, _ := json.Marshal(message)
	res := goWeChat.PostJson(url, string(req))
	textResponse := model.TextMessageResponse{}
	json.Unmarshal([]byte(res), &textResponse)
	return textResponse
}

func (goWeChat *GoWeChat) SendMarkDownMessage(message model.MarkdownMessage) model.TextMessageResponse {
	url := goWeChat.Url + "/cgi-bin/message/send?access_token=" + goWeChat.Access_token
	req, _ := json.Marshal(message)
	res := goWeChat.PostJson(url, string(req))
	textResponse := model.TextMessageResponse{}
	json.Unmarshal([]byte(res), &textResponse)
	return textResponse
}

func (goWeChat *GoWeChat) Get(url string) string {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		base := model.Base{}
		base.Errcode = -1
		base.Errmsg = err.Error()
		re, _ := json.Marshal(base)
		return string(re)
	}
	return string(resp.Body())
}

func (goWeChat *GoWeChat) PostJson(url string, body string) string {
	client := resty.New()
	resp, err := client.R().
		SetBody(body).
		Post(url)
	if err != nil {
		base := model.Base{}
		base.Errcode = -1
		base.Errmsg = err.Error()
		re, _ := json.Marshal(base)
		return string(re)
	} else {
		return string(resp.Body())
	}
}
