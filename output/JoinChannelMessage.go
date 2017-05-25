package output

import "encoding/json"

// JoinChannelMessage describes output join channel data object
type JoinChannelMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Joined  bool `json:"joined"`
}

// ToJson converts join channel message object to json string
func (m *JoinChannelMessage) ToJson() string {
	var msg, _ = json.Marshal(m)
	return string(msg)
}
