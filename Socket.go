package socket

import (
	"golang.org/x/net/websocket"
	"log"
	"github.com/thedarkphoton/socket/input"
	"github.com/thedarkphoton/socket/output"
)

// Channel defines all different types of channels
type Channel int

const (
	// ALL channel
	ALL Channel = iota
)

var sock *Socket

// Connections maintains websocket connections alive
func Connections(ws *websocket.Conn) {
	id := string(len(sock.channels))
	id = "10"
	// log.Println("test" + id)
	sock.connect(id, ws)
	sock.connected(id)

	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println("Unable to receive")
		}

		json := input.ToMessage(&msg)
		switch json.Type {
		case "IPingMessage":
			sock.ping(input.ToPingMessage(&msg))
		}
	}
}

// Socket represents all connections to the socket
type Socket struct {
	channels map[Channel](map[string]*websocket.Conn)
}

// Instance gets instance of a socket
func Instance() *Socket {
	sock = &Socket{channels: make(map[Channel](map[string]*websocket.Conn))}
	sock.channels[ALL] = make(map[string]*websocket.Conn)
	return sock
}

// Connect processes connection messages
func (s *Socket) connect(id string, ws *websocket.Conn) {
	s.channels[ALL][id] = ws
}

func (s *Socket) connected(id string) {
	var out = output.ConnectedMessage{Type:"OConnectedMessage", Id: id}
	log.Println(out.ToJson())
	if err := websocket.Message.Send(s.channels[ALL][id], out.ToJson()); err != nil {
		log.Println("Unable to send")
	}
}

func (s *Socket) ping(msg input.PingMessage) {
	var out = output.PingMessage{Type:"OPingMessage", Start:msg.Start}
	log.Println(msg)
	if err := websocket.Message.Send(s.channels[ALL][msg.Id], out.ToJson()); err != nil {
		log.Println("Unable to send")
	}
}

// JoinChannel adds connection to the specified channel
func (s *Socket) joinChannel(name Channel, id string, ws *websocket.Conn) {
	s.channels[name][id] = ws
}

// ConnectionsInChannel gets all connections in the specified channel
func (s Socket) ConnectionsInChannel(name Channel) []Connection {
	return toConnections(s.channels[name])
}
