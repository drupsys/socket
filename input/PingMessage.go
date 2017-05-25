package input

import "encoding/json"

// PingMessage describes input ping data object
type PingMessage struct {
	Start int `json:"start"`
}

// ToMessage converts json string to message object
func ToPingMessage(raw *string) PingMessage {
	var data = PingMessage{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}
