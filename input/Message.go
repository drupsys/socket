package input

import (
	"encoding/json"
)

// Message describes input message data object
type Message struct {
	Type string `json:"type"`
}

// ToMessage converts json string to message object
func ToMessage(raw *string) Message {
	var data = Message{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

// PingMessage describes input ping data object
type PingMessage struct {
	Id string `json:"id"`
	Start int `json:"start"`
}

// ToMessage converts json string to message object
func ToPingMessage(raw *string) PingMessage {
	var data = PingMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

// JoinChannelMessage describes input join channel data object
type JoinChannelMessage struct {
	Id string `json:"id"`
	Channel string `json:"channel"`
}

// ToJoinChannelMessage converts json string to join channel message object
func ToJoinChannelMessage(raw *string) JoinChannelMessage {
	var data = JoinChannelMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

// ChannelMessage describes input channel data object
type ChannelMessage struct {
	Id string `json:"id"`
	Channel string `json:"channel"`
	Data string `json:"data"`
}

// ToChannelMessage converts json string to channel message object
func ToChannelMessage(raw *string) ChannelMessage {
	var data = ChannelMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}

// SendMessage describes input send data object
type SendMessage struct {
	Id string `json:"id"`
	Data string `json:"data"`
}

// ToSendMessage converts json string to send message object
func ToSendMessage(raw *string) SendMessage {
	var data = SendMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}