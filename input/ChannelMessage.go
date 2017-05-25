package input

import "encoding/json"

// ChannelMessage describes input channel data object
type ChannelMessage struct {
	Channel string `json:"channel"`
	Data string `json:"data"`
}

// ToChannelMessage converts json string to channel message object
func ToChannelMessage(raw *string) ChannelMessage {
	var data = ChannelMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}
