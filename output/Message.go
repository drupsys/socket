package output

type Type int
const (
	CONNECTED Type = iota
	PING Type = iota
	LOOPBACK Type = iota
	JOIN_CHANNEL Type = iota
	CHANNEL Type = iota
)
