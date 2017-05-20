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