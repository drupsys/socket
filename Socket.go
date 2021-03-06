package socket

import (
	"golang.org/x/net/websocket"
	"log"
	"github.com/thedarkphoton/socket/input"
	"github.com/thedarkphoton/socket/output"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// sock static instance of a web socket
var sock *WebSocket

// randomBytes generates random bytes
func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

// randomBase64 generates random base 64 string
func randomBase64(n int) (string, error) {
	b, err := randomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

// Connections maintains websocket connections alive
func Connections(ws *websocket.Conn) {
	id, err := sock.connect(ws)
	if err != nil {
		log.Printf("Connection could not be established, failed with: %v", err)
		return
	} else {
		log.Println("Connected user with id: " + id)
	}
	
	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Printf("Unable to receive, failed with: %v", err)
			sock.disconnect(id)
			return
		}

		json := input.ToMessage(&msg)
		switch json.Type {
		case "PingMessage":
			if err := sock.ping(id, input.ToPingMessage(&msg)); err != nil {
				log.Printf("Unble to respond with Ping message, failed with: %v", err)
			}
		case "JoinChannelMessage":
			if err := sock.join(id, input.ToJoinChannelMessage(&msg)); err != nil {
				log.Printf("Unble to respond with JoinChannel message, failed with: %v", err)
			}
		case "ChannelMessage":
			if err := sock.sendToChannel(id, input.ToChannelMessage(&msg)); err != nil {
				log.Printf("Unble to respond with Channel message, failed with: %v", err)
			}
		case "LoopbackMessage":
			if err := sock.sendToConnection(id, input.ToSendMessage(&msg)); err != nil {
				log.Printf("Unble to respond with Send message, failed with: %v", err)
			}
		}
	}
}

// Instance gets instance of a socket
func Instance() *WebSocket {
	sock = &WebSocket{
		connections: make(map[string]Connection),
		channels:    make(map[string]Channel)}
	sock.connections["all"] = make(Connection)
	sock.channels["all"] = func(id, data string) string {
		return "Loopback callback not implemented"
	}
	
	return sock
}

// Channel describes the callback function of the channel
type Channel (func(id, inJson string) string)

// Connection describes the connection type
type Connection (map[string]*websocket.Conn)

// WebSocket describes data of all web socket's connections
type WebSocket struct {
	connections map[string]Connection
	channels    map[string]Channel
}

// SetContextChannel sets the callback for the LoopbackMessage action
func (s *WebSocket) SetContextChannel(callback Channel) {
	sock.channels["all"] = callback
}

// AddChannel creates a channel and associates callback function for the ChannelMessage action to it
func (s *WebSocket) AddChannel(channel string, callback Channel) {
	if channel == "all" {
		return
	}
	
	sock.connections[channel] = make(Connection)
	sock.channels[channel] = callback
}

// RemoveChannel removes channel and its associated function
func (s *WebSocket) RemoveChannel(channel string) {
	sock.connections[channel] = nil
	sock.channels[channel] = nil
}

// connect creates a client id and send it back to the client
func (s *WebSocket) connect(ws *websocket.Conn) (string, error) {
	var id string
	for loop := true; loop; loop = s.connections["all"][id] != nil {
		if tmp, err := randomBase64(16); err != nil {
			return "", err
		} else {
			id = tmp
		}
	}
	
	s.connections["all"][id] = ws
	var out = output.ConnectedMessage{Type:"ConnectedMessage", Id: id}
	if err := websocket.Message.Send(s.connections["all"][id], out.ToJson()); err != nil {
		return "", err
	}
	
	return id, nil
}

// disconnect removes all client's data
func (s *WebSocket) disconnect(id string) {
	for channel := range s.connections {
		s.connections[channel][id] = nil
	}
}

// ping sends a ping response to the client
func (s *WebSocket) ping(id string, msg input.PingMessage) error {
	var out = output.PingMessage{Type:"PingMessage", Start:msg.Start}
	if ws := s.connections["all"][id]; ws != nil {
		if err := websocket.Message.Send(ws, out.ToJson()); err != nil {
			return err
		}
	} else {
		return errors.New("Web socket with the specified id was not found")
	}
	
	return nil
}

// join adds connection to the specified channel
func (s *WebSocket) join(id string, msg input.JoinChannelMessage) error {
	var err error
	if s.channels[msg.Channel] == nil {
		err = errors.New("User tried to join non-existing channel")
	} else if s.connections["all"][id] == nil {
		err = errors.New("Invalid user tried to join a channel")
	} else {
		s.connections[msg.Channel][id] = s.connections["all"][id]
	}
	
	var out = output.JoinChannelMessage{Type: "JoinChannelMessage", Joined: err == nil}
	if ws := s.connections["all"][id]; ws != nil {
		if err := websocket.Message.Send(ws, out.ToJson()); err != nil {
			return err
		}
	} else {
		return errors.New("Web socket with the specified id was not found")
	}
	
	return err
}

// sendToChannel receives data from a client, passes it to the channel's callback and send back
// the result of the callback to all client of the channel (except sender)
func (s *WebSocket) sendToChannel(id string, msg input.ChannelMessage) error {
	if _, ok := s.channels[msg.Channel]; !ok {
		return errors.New("User sent data to non-existing channel")
	}
	
	var err error
	var out = output.ChannelMessage{Type:"ChannelMessage", Channel: msg.Channel, Data: s.channels[msg.Channel](id, msg.Data)}
	for _, ws := range s.connections[msg.Channel] {
		if ws == nil {
			continue
		}
		
		if tmpErr := websocket.Message.Send(ws, out.ToJson()); err != nil {
			err = tmpErr
		}
	}
	
	return err
}

// sendToChannel receives data from a client, passes it to the context's callback and send back
// the result of the callback to the sender
func (s *WebSocket) sendToConnection(id string, msg input.LoopbackMessage) error {
	if s.connections["all"][id] == nil {
		return errors.New("Invalid user tried to send data")
	}
	
	var out = output.LoopbackMessage{Type: "LoopbackMessage", Data: s.channels["all"](id, msg.Data)}
	if err := websocket.Message.Send(s.connections["all"][id], out.ToJson()); err != nil {
		return err
	}
	
	return nil
}