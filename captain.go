package main

import (
	"net/http"
	"encoding/json"
	"bytes"
)

type Captain struct{}

type DingTalkMsg struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}
type Text struct {
	Content string `json:"content"`
}

func (ct *Captain) Dispose(corsairInfo *CorsairInfo) (err error) {
	dtm := DingTalkMsg{
		MsgType: "text",
		Text: Text{
			Content: corsairInfo.Msg,
		},
	}
	dtmBytes, err := json.Marshal(dtm)

	_, err = http.Post("https://oapi.dingtalk.com/robot/send?access_token=c5d88af08d421bab44020fad40d58bf8137a258959716d9c736b9e730b2bc551", "application/json", bytes.NewBuffer(dtmBytes))
	return
}
