package protocol

type Protocol int

const (
	HTTP Protocol = iota
	HTTPS
	TCP
	WEBSOCKET
	UDP
	IPv4
)

const (
	_http      = "http"
	_https     = "https"
	_tcp       = "tcp"
	_udp       = "udp"
	_websocket = "websocket"
	_ipv4      = "ipv4"
)

func New(in string) Protocol {
	switch in {
	case _http:
		return HTTP
	case _https:
		return HTTPS
	case _tcp:
		return TCP
	case _udp:
		return UDP
	case _websocket:
		return WEBSOCKET
	case _ipv4:
		return IPv4
	default:
		return TCP
	}
}

func (p Protocol) String() string {
	switch p {
	case HTTP:
		return _http
	case HTTPS:
		return _https
	case TCP:
		return _tcp
	case UDP:
		return _udp
	case WEBSOCKET:
		return _websocket
	case IPv4:
		return _ipv4
	default:
		return _tcp
	}
}
