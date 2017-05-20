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

var sock *Socket

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

func defaultContextCallback(id, data string) string {
	return "Context callback not implemented"
}

// Connections maintains websocket connections alive
func Connections(ws *websocket.Conn) {
	id, err := sock.connect(ws)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Connected user with id: " + id)
	}
	
	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println(err)
		}

		json := input.ToMessage(&msg)
		
		switch json.Type {
		case "IPingMessage":
			if err := sock.ping(input.ToPingMessage(&msg)); err != nil {
				log.Println(err)
			}
		case "IJoinChannelMessage":
			if err := sock.join(input.ToJoinChannelMessage(&msg)); err != nil {
				log.Println(err)
			}
		case "IChannelMessage":
			if err := sock.sendToChannel(input.ToChannelMessage(&msg)); err != nil {
				log.Println(err)
			}
		case "ISentMessage":
			if err := sock.sendToConnection(input.ToSendMessage(&msg)); err != nil {
				log.Println(err)
			}
		}
	}
}

// Instance gets instance of a socket
func Instance() *Socket {
	sock = &Socket{
		connections: make(map[string]Connection),
		channels:    make(map[string]Channel)}
	sock.connections["all"] = make(Connection)
	sock.channels["all"] = defaultContextCallback
	return sock
}

type Channel (func(id, inJson string) string)
type Connection (map[string]*websocket.Conn)

// Socket represents all connections to the socket
type Socket struct {
	connections map[string]Connection
	channels    map[string]Channel
}

func (s *Socket) SetContextChannel(callback Channel) {
	sock.channels["all"] = callback
}

func (s *Socket) CreateChannel(channel string, callback Channel) {
	if channel == "all" {
		return
	}
	
	sock.connections[channel] = make(Connection)
	sock.channels[channel] = callback
}

func (s *Socket) RemoveChannel(channel string) {
	sock.connections[channel] = nil
	sock.channels[channel] = nil
}

// Connect processes connection messages
func (s *Socket) connect(ws *websocket.Conn) (string, error) {
	var id string
	for loop := true; loop; loop = s.connections["all"][id] != nil {
		if tmp, err := randomBase64(16); err != nil {
			return "", err
		} else {
			id = tmp
		}
	}
	
	s.connections["all"][id] = ws
	var out = output.ConnectedMessage{Type:"OConnectedMessage", Id: id}
	if err := websocket.Message.Send(s.connections["all"][id], out.ToJson()); err != nil {
		return "", err
	}
	
	return id, nil
}

// ping sends a ping response to the client
func (s *Socket) ping(msg input.PingMessage) error {
	var out = output.PingMessage{Type:"OPingMessage", Start:msg.Start}
	if ws := s.connections["all"][msg.Id]; ws != nil {
		if err := websocket.Message.Send(ws, out.ToJson()); err != nil {
			return err
		}
	} else {
		return errors.New("Web socket with the specified id was not found")
	}
	
	return nil
}

// join adds connection to the specified channel
func (s *Socket) join(msg input.JoinChannelMessage) error {
	var err error
	if s.channels[msg.Channel] == nil {
		err = errors.New("User tried to join non-existing channel")
	} else if s.connections["all"][msg.Id] == nil {
		err = errors.New("Invalid user tried to join a channel")
	} else {
		s.connections[msg.Channel][msg.Id] = s.connections["all"][msg.Id]
	}
	
	var out = output.JoinChannelMessage{Type:"OJoinChannelMessage", Status:err == nil}
	if ws := s.connections["all"][msg.Id]; ws != nil {
		if err := websocket.Message.Send(ws, out.ToJson()); err != nil {
			return err
		}
	} else {
		return errors.New("Web socket with the specified id was not found")
	}
	
	return err
}

func (s *Socket) sendToChannel(msg input.ChannelMessage) error {
	if s.channels[msg.Channel] == nil {
		return errors.New("User sent data to non-existing channel")
	}
	
	var err error
	var out = output.ChannelMessage{Type:"OChannelMessage", Data: s.channels[msg.Channel](msg.Id, msg.Data)}
	for id, ws := range s.connections[msg.Channel] {
		if id == msg.Id {
			continue
		}
		
		if tmpErr := websocket.Message.Send(ws, out.ToJson()); err != nil {
			err = tmpErr
		}
	}
	
	return err
}

func (s *Socket) sendToConnection(msg input.SendMessage) error {
	if s.connections["all"][msg.Id] == nil {
		return errors.New("Invalid user tried to send data")
	}
	
	var out = output.SendMessage{Type:"OChannelMessage", Data: s.channels["all"](msg.Id, msg.Data)}
	if err := websocket.Message.Send(s.connections["all"][msg.Id], out.ToJson()); err != nil {
		return err
	}
	
	return nil
}