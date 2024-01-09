package go_libp2p_grpc

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"net"
)

var _ net.Conn = (*Conn)(nil)

type Conn struct {
	network.Stream
}

// LocalAddr returns the local address.
func (c *Conn) LocalAddr() net.Addr {
	return toNetAddr(c.Stream.Conn().LocalMultiaddr())
}

// RemoteAddr returns the remote address.
func (c *Conn) RemoteAddr() net.Addr {
	return toNetAddr(c.Stream.Conn().RemoteMultiaddr())
}

func toNetAddr(ma multiaddr.Multiaddr) net.Addr {
	na, err := manet.ToNetAddr(ma)
	if err != nil {
		return fakeAddr()
	}
	return na
}
