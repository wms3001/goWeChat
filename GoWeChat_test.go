package goWeChat

import (
	"fmt"
	"github.com/wms3001/goWeChat/model"
	"log"
	"testing"
)

var goWeChat GoWeChat

func TestGoWeChat_GetAccessToken(t *testing.T) {
	goWeChat.Url = "https://qyapi.weixin.qq.com"
	goWeChat.Corpid = "wwcfcac37088fc0f2433"
	goWeChat.Corpsecret = "Wo33GJC6t4nuVjgRsCxyM0ArhiKaJ8UPk6v_SZYavh0Tc"
	token := goWeChat.GetAccessToken()
	log.Println(token.AccessToken)
}

func TestGoWeChat_SendTextMessage(t *testing.T) {
	message := model.TextMessage{}
	message.Touser = "qy01669528189220079ce2e58dd8"
	message.Agentid = 1000005
	message.Msgtype = "text"
	message.Safe = 0
	message.EnableIdTrans = 0
	message.EnableDuplicateCheck = 0
	message.DuplicateCheckInterval = 1000
	message.Totag = ""
	message.Toparty = ""
	message.Text.Content = "This is a Test!"
	res := goWeChat.SendTextMessage(message)
	log.Println(fmt.Sprint(res.Errcode) + res.Errmsg)
}
