# golang推送企业微信应用消息

------

实现了企业微信各种应用推送消接口

------

1. 安装
```go
go get github.com/wms3001/goWeChat
```
2. 获取access_token
```go
goWeChat.Url = "https://qyapi.weixin.qq.com"
goWeChat.Corpid = "wwcfcac37088fc0f2433"
goWeChat.Corpsecret = "Wo33GJC6t4nuVjgRsCxyM0ArhiKaJ8UPk6v_SZYavh0Tc"
token := goWeChat.GetAccessToken()
log.Println(token.AccessToken)
```
3. 发送文本消息
```go
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
```