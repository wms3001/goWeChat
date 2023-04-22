package model

type TextMessage struct {
	MessageBase
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

type TextMessageResponse struct {
	Base
	Msgid string `json:"msgid"`
}
