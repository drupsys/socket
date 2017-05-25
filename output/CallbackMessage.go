package output

import "encoding/json"

// CallbackMessage describes output ping data object
type CallbackMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// ToJson converts send message object to json string
func (m *CallbackMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}