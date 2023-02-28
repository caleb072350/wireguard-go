package device

import "net"

const (
	ConnRoutineNumber = 2
)

/* A Bind handles listening on a port for both IPv6 and IPv4 UDP traffic
 */
type Bind interface {
	SetMark(value uint32) error
	ReceiveIPv6(buff []byte) (int, Endpoint, error)
	ReceiveIPv4(buff []byte) (int, Endpoint, error)
	Send(buff []byte, end Endpoint) error
	Close() error
}

/* An Endpoint maintains the source/destination catching for a peer
 *
 * dst : the remote address of a peer ("endpoint" in uapi terminology)
 * src : the local address from which datagrams originate going to the peer
 */
type Endpoint interface {
	ClearSrc()           // clears the source address
	SrcToString() string // returns the local source address (ip:port)
	DstToString() string // returns the destination address (ip:port)
	DstToBytes() []byte  // used for mac2 cookie calculations
	DstIP() net.IP
	SrcIP() net.IP
}
