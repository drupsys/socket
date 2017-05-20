package output

import "encoding/json"

type PingMessage struct {
	Type string `json:"type"`
	Start int `json:"start"`
}

func (m *PingMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

type ConnectedMessage struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

func (m *ConnectedMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

type JoinChannelMessage struct {
	Type string `json:"type"`
	Status bool `json:"status"`
}

func (m *JoinChannelMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

type ChannelMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (m *ChannelMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

type SendMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (m *SendMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}