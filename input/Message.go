package input

import (
	"encoding/json"
)

type Message struct {
	Type string `json:"type"`
}

func ToMessage(raw *string) Message {
	var data = Message{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

type PingMessage struct {
	Id string `json:"id"`
	Start int `json:"start"`
}

func ToPingMessage(raw *string) PingMessage {
	var data = PingMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

type JoinChannelMessage struct {
	Id string `json:"id"`
	Channel string `json:"channel"`
}

func ToJoinChannelMessage(raw *string) JoinChannelMessage {
	var data = JoinChannelMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

type ChannelMessage struct {
	Id string `json:"id"`
	Channel string `json:"channel"`
	Data string `json:"data"`
}

func ToChannelMessage(raw *string) ChannelMessage {
	var data = ChannelMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

type SendMessage struct {
	//Type string `json:"type"`
	Id string `json:"id"`
	Data string `json:"data"`
}

func ToSendMessage(raw *string) SendMessage {
	var data = SendMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}