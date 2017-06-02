package output

import "encoding/json"

// LoopbackMessage describes output ping data object
type LoopbackMessage struct {
	Type Type `json:"type"`
	Data string `json:"data"`
}

// ToJson converts send message object to json string
func (m *LoopbackMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}