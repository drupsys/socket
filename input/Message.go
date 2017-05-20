package input

import (
	"encoding/json"
)

type Message struct {
	Type string `json:"type"`
}

type PingMessage struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Start int `json:"start"`
}

func ToMessage(raw *string) Message {
	var data = Message{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

func ToPingMessage(raw *string) PingMessage {
	var data = PingMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}