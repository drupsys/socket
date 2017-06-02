package input

import (
	"encoding/json"
)

type Type int
const (
	PING Type = iota
	LOOPBACK Type = iota
	JOIN_CHANNEL Type = iota
	CHANNEL Type = iota
)

// Message describes input message data object
type Message struct {
	Type Type `json:"type"`
}

// ToMessage converts json string to message object
func ToMessage(raw *string) Message {
	var data = Message{}
	json.Unmarshal([]byte(*raw), &data)
	return data
}