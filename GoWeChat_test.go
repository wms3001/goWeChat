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
	goWeChat.Corpid = "wwc33fcac37088fc0f24"
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

func TestGoWeChat_SendMarkDownMessage(t *testing.T) {
	message := model.MarkdownMessage{}
	message.Touser = "qy01669528189220079ce2e58dd8"
	message.Agentid = 1000005
	message.Msgtype = "markdown"
	message.Safe = 0
	message.EnableIdTrans = 0
	message.EnableDuplicateCheck = 0
	message.DuplicateCheckInterval = 1000
	message.Totag = ""
	message.Toparty = ""
	message.Markdown.Content = "# 您的会议室已经预定，稍后会同步到`邮箱` \n                                >**事项详情** \n                                >事　项：<font color=\"info\">开会</font> \n                                >组织者：@miglioguan \n                                >参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang \n                                > \n                                >会议室：<font color=\"info\">广州TIT 1楼 301</font> \n                                >日　期：<font color=\"warning\">2018年5月18日</font> \n                                >时　间：<font color=\"comment\">上午9:00-11:00</font> \n                                > \n                                >请准时参加会议。 \n                                > \n                                >如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"
	res := goWeChat.SendMarkDownMessage(message)
	log.Println(fmt.Sprint(res.Errcode) + res.Errmsg)
}
