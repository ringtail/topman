package seaway

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
)

type Captain struct{
	PhoneNumber string
}

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
	fmt.Printf("Yes,commander. Enemies are found: %s\n",corsairInfo.Msg)
	_, err = http.Post(ct.PhoneNumber, "application/json", bytes.NewBuffer(dtmBytes))
	return
}
