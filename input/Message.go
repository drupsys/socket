package input

import (
	"encoding/json"
)

// Message describes input message data object
type Message struct {
	Type string `json:"type"`
}

// ToMessage converts json string to message object
func ToMessage(raw *string) Message {
	var data = Message{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}