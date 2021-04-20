package protocol

type Protocol int

const (
	HTTP Protocol = iota
	HTTPS
	TCP
	WEBSOCKET
	UDP
	IPv4
	GRPC
)

const (
	_http      = "http"
	_https     = "https"
	_tcp       = "tcp"
	_udp       = "udp"
	_websocket = "websocket"
	_ipv4      = "ipv4"
	_grpc      = "grpc"
)

func (p Protocol) String() string {
	return [...]string{_http, _https, _tcp, _websocket, _udp, _ipv4, _grpc}[p]
}

