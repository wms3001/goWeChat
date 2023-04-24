package model

type MarkdownMessage struct {
	MessageBase
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}
