package output

import "encoding/json"

// ChannelMessage describes output ping data object
type ChannelMessage struct {
	Type Type `json:"type"`
	Channel string `json:"channel"`
	Data string `json:"data"`
}

// ToJson converts channel message object to json string
func (m *ChannelMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}