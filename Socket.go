package socket

import (
	"golang.org/x/net/websocket"
)

// Socket describes
type Socket struct {
	channels map[string](map[string]*websocket.Conn)
}

// JoinChannel adds connection to the specified channel
func (s *Socket) JoinChannel(name, id string, ws *websocket.Conn) {
	s.channels[name][id] = ws
}

// ConnectionsInChannel asd
func (s Socket) ConnectionsInChannel(name string) []Connection {
	return toConnections(s.channels[name])
}
