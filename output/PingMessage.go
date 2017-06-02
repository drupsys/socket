package output

import "encoding/json"

// PingMessage describes output ping data object
type PingMessage struct {
	Type Type `json:"type"`
	Start int `json:"start"`
}

// ToJson converts ping message object to json string
func (m *PingMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}