package output

import "encoding/json"

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