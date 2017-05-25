package input

import "encoding/json"

// LoopbackMessage describes input send data object
type LoopbackMessage struct {
	Data string `json:"data"`
}

// ToSendMessage converts json string to send message object
func ToSendMessage(raw *string) LoopbackMessage {
	var data = LoopbackMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}
