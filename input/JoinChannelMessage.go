package input

import "encoding/json"

// JoinChannelMessage describes input join channel data object
type JoinChannelMessage struct {
	Channel string `json:"channel"`
}

// ToJoinChannelMessage converts json string to join channel message object
func ToJoinChannelMessage(raw *string) JoinChannelMessage {
	var data = JoinChannelMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}
