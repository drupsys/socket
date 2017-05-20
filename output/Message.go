package output

import "encoding/json"

// PingMessage describes output ping data object
type PingMessage struct {
	Type string `json:"type"`
	Start int `json:"start"`
}

// ToJson converts ping message object to json string
func (m *PingMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

// ConnectedMessage describes output ping data object
type ConnectedMessage struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// ToJson converts connected message object to json string
func (m *ConnectedMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

// JoinChannelMessage describes output ping data object
type JoinChannelMessage struct {
	Type string `json:"type"`
	Status bool `json:"status"`
}

// ToJson converts join channel message object to json string
func (m *JoinChannelMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

// ChannelMessage describes output ping data object
type ChannelMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// ToJson converts channel message object to json string
func (m *ChannelMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}

// SendMessage describes output ping data object
type SendMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// ToJson converts send message object to json string
func (m *SendMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}