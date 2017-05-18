package socket

import (
	"golang.org/x/net/websocket"
)

// Connection represents a socket connection
type Connection struct {
	ID string
	WS *websocket.Conn
}

// ToConnections convert the map of websocket connections to array of connections
func toConnections(conn map[string]*websocket.Conn) []Connection {
	var result []Connection
	for k, v := range conn {
		result = append(result, Connection{ID: k, WS: v})
	}

	return result
}
